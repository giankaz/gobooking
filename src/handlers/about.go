package handlers

import (
	"net/http"

	"github.com/giankaz/gobooking/src/models"
	"github.com/giankaz/gobooking/src/render"
)

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{})
}
