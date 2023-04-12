package render

import (
	"net/http"
	"testing"

	"github.com/giankaz/gobooking/src/models"
)

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "../templates"
	_, err := CreateTemplateCache()

	if err != nil {
		t.Error(err)
	}

}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestAddDefaultData(t *testing.T) {
	var templateData models.TemplateData

	req, err := getSession()

	if err != nil {
		t.Error(err)
	}

	session.Put(req.Context(), "flash", "123")

	result := AddDefaultData(&templateData, req)

	if result.Flash != "123" {
		t.Error("flash value is not 123 it is: ", result.Flash)
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "../templates"
	templateCache, err := CreateTemplateCache()

	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = templateCache

	req, _ := getSession()

	var ww mockWriter

	err = RenderTemplate(&ww, req, "home.page.html", &models.TemplateData{})

	if err != nil {
		t.Error(err)
	}

	err = RenderTemplate(&ww, req, "fakefakefake.page.html", &models.TemplateData{})

	if err == nil {
		t.Error("rendered template that doesnt exist")
	}
}

func getSession() (*http.Request, error) {
	req, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		return nil, err
	}

	context := req.Context()

	context, _ = session.Load(context, req.Header.Get("X-Session"))

	req = req.WithContext(context)

	return req, nil

}

type mockWriter struct {
}

func (mw *mockWriter) Header() http.Header {
	var header http.Header
	return header
}

func (mw *mockWriter) Write(sliceofbytes []byte) (int, error) {
	return len(sliceofbytes), nil
}

func (mw *mockWriter) WriteHeader(statuscode int) {

}
