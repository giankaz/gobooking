package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/giankaz/gobooking/src/config"
	"github.com/giankaz/gobooking/src/models"
	"github.com/giankaz/gobooking/src/render"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
)

var app config.AppConfig

var testServer *httptest.Server

var session *scs.SessionManager

var infoLog *log.Logger
var errorLog *log.Logger

var pathToTemplates = "../templates"

func TestMain(m *testing.M) {
	routes := getRoutes()

	testServer = httptest.NewTLSServer(routes)

	defer testServer.Close()

	os.Exit(m.Run())
}

func getRoutes() http.Handler {
	//registrar structs na sessão
	gob.Register(models.Reservation{})

	session = scs.New()

	infoLog = log.New(os.Stdout, "[INFO - BOOKINGS]:\t", log.Ldate|log.Ltime)

	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "[ERROR - BOOKINGS]:\t", log.Ldate|log.Ltime|log.Lshortfile)

	app.ErrorLog = errorLog

	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true

	session.Cookie.SameSite = http.SameSiteLaxMode

	session.Cookie.Secure = false

	app.Session = session

	templateCache, err := CreateTestTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	//change this to true in test
	app.UseCache = true

	repo := NewRepository(&app)

	NewHandlers(repo)

	render.NewTemplates(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	// cpomment for test
	//mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)

	mux.Get("/generals-quarters", Repo.Generals)

	mux.Get("/majors-suite", Repo.Majors)

	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Get("/search-availability-json", Repo.AvailabilityJSON)

	mux.Get("/contact", Repo.Contact)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)
	mux.Post("/make-reservation", Repo.PostReservation)

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {
	log.Println("Start of Complex Renderer")

	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))

	log.Println(pages)

	if err != nil {
		log.Println(err)

		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.New(name).ParseFiles(page)

		if err != nil {
			log.Println(err)

			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))

		if err != nil {
			log.Println(err)

			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))

			if err != nil {
				log.Println("End of Complex Renderer")

				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	log.Println("End of Complex Renderer")

	return myCache, nil

}

func getTestCase(t *testing.T, path string) {
	resp, err := testServer.Client().Get(testServer.URL + path)

	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expect status code %d, instead got %d", 200, resp.StatusCode)
	}
}

func postTestCase(t *testing.T, path string, body map[string]string) {
	values := url.Values{}

	for key, value := range body {
		log.Println(key, value)
		values.Add(key, value)
	}

	resp, err := testServer.Client().PostForm(testServer.URL+path, values)

	if body["EXPECT_ERROR"] == "EXPECT" {
		if resp.StatusCode == 200 {
			log.Println(resp)
			t.Errorf("not expect status code %d, instead got %d", 200, resp.StatusCode)
		}
	} else {
		if err != nil {
			t.Log(err)
			t.Fatal(err)
		}
	}

	if resp.StatusCode != 200 {
		t.Errorf("expect status code %d, instead got %d", 200, resp.StatusCode)
	}
}
