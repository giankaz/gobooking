package models

import "github.com/giankaz/gobooking/src/forms"

// TemplateData Holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float64
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
	Form *forms.Form
}