package handlers

import (
	"net/http"

	"github.com/anubhavjoshi040/go-webapp-template/config"
	"github.com/anubhavjoshi040/go-webapp-template/pkg/models"
	"github.com/anubhavjoshi040/go-webapp-template/pkg/render"
)

// Repository is the repository type
type Repository struct{
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}



