package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/giankaz/gobooking/src/models"
	"github.com/giankaz/gobooking/src/render"
)

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintln("Posted to search availability", start, end)))
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse{
		Ok:      true,
		Message: "available!",
	}

	output, err := json.MarshalIndent(response, "", "  ")

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
