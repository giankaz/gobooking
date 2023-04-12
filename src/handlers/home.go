package handlers

import (
	"net/http"

	"github.com/giankaz/gobooking/src/models"
	"github.com/giankaz/gobooking/src/render"
)

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}
