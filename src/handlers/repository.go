package handlers

import "github.com/giankaz/gobooking/src/config"

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepository creates a new repository
func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(repo *Repository) {
	Repo = repo
}
