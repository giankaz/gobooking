package main

import (
	"testing"

	"github.com/giankaz/gobooking/src/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch typeof := mux.(type) {
	case *chi.Mux:
		// passed
	default:
		t.Errorf("type is not chi.mux, time is %T", typeof)
	}
}
