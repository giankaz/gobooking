package handlers

import (
	"net/http"

	"github.com/giankaz/gobooking/src/models"
	"github.com/giankaz/gobooking/src/render"
)

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}
