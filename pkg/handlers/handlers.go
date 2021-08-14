package handlers

import (
	"net/http"

	"github.com/corentinclichy/myapp/pkg/config"
	"github.com/corentinclichy/myapp/pkg/models"
	"github.com/corentinclichy/myapp/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// New repo create a new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

//NewHandler set the repository for the new handler
func NewHandler(r *Repository) {
	Repo = r
}

// @Title Home
// @Description this is the home Page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// @Title about
// @Description this is the about Page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	stringMap["test"] = "Hello world"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
