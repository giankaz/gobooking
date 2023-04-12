package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/giankaz/gobooking/src/config"
	"github.com/giankaz/gobooking/src/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig
var pathToTemplates = "./src/templates"

func NewTemplates(conf *config.AppConfig) {
	app = conf
}

func AddDefaultData(templateData *models.TemplateData, req *http.Request) *models.TemplateData {
	templateData.Flash = app.Session.PopString(req.Context(), "flash")
	templateData.Error = app.Session.PopString(req.Context(), "error")
	templateData.Warning = app.Session.PopString(req.Context(), "warning")
	templateData.CSRFToken = nosurf.Token(req)
	return templateData
}

// ComplexRenderTemplate renders a template to a response
func RenderTemplate(w http.ResponseWriter, req *http.Request, tmpl string, data *models.TemplateData) error {
	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		tc, err := CreateTemplateCache()
		if err != nil {
			return errors.New(err.Error())

		}
		templateCache = tc
	}

	template, templateExists := templateCache[tmpl]

	if !templateExists {
		return errors.New("missing template")
	}

	buffer := new(bytes.Buffer)

	data = AddDefaultData(data, req)

	templateError := template.Execute(buffer, data)

	if templateError != nil {
		log.Fatal(templateError)

	}

	_, bufferError := buffer.WriteTo(w)

	if bufferError != nil {
		return errors.New(bufferError.Error())
	}

	return nil

}

func CreateTemplateCache() (map[string]*template.Template, error) {
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
