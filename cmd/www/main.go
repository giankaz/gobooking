package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/giankaz/gobooking/src/config"
	"github.com/giankaz/gobooking/src/handlers"
	"github.com/giankaz/gobooking/src/models"
	"github.com/giankaz/gobooking/src/render"
	"github.com/giankaz/gobooking/src/helpers"

	"github.com/alexedwards/scs/v2"
)

const PORT = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := run()

	if err != nil {
		log.Println(err)
	}

	server := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//registrar structs na sessão
	gob.Register(models.Reservation{})

	infoLog = log.New(os.Stdout, "[INFO - BOOKINGS]:\t", log.Ldate|log.Ltime)

	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "[ERROR - BOOKINGS]:\t", log.Ldate|log.Ltime|log.Lshortfile)

	app.ErrorLog = errorLog

	session = scs.New()

	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true

	session.Cookie.SameSite = http.SameSiteLaxMode

	session.Cookie.Secure = false

	app.Session = session

	templateCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = templateCache
	//change this to true in prod
	app.UseCache = false

	repo := handlers.NewRepository(&app)

	helpers.NewHelpers(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	log.Printf("Running on PORT %v, Gratz!", PORT)

	return nil
}
