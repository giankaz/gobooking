package handlers

import (
	"net/http"

	"github.com/giankaz/gobooking/src/models"
	"github.com/giankaz/gobooking/src/render"
)

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}
