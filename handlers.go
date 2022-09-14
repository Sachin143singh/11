package handlers

import (
	"net/http"

	"github.com/Sachin47singh/go-course/pkg/config"
	"github.com/Sachin47singh/go-course/pkg/models"
	"github.com/Sachin47singh/go-course/pkg/render"
)

// // // Home is the handler for the home page
// func Home(w http.ResponseWriter, r *http.Request) {
// 	render.RenderTemplate(w, "home.page.tmpl")
// }

// // // About is the handler for the about page
// func About(w http.ResponseWriter, r *http.Request) {
// 	render.RenderTemplate(w, "about.page.tmpl")
// }

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

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

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
