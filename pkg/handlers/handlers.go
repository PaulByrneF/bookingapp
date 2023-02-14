package handlers

import (
	"net/http"

	"github.com/paulbyrne/bookingapp/pkg/config"
	"github.com/paulbyrne/bookingapp/pkg/models"
	"github.com/paulbyrne/bookingapp/pkg/render"
)

// Repository Design Pattern
// Repo the repository used by handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	AppConfig *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {

	// retrieve ip address from http request and store in appConfig.session
	remoteIP := r.RemoteAddr
	repo.AppConfig.Session.Put(r.Context(), "remote_ip", remoteIP)

	// render template & inject datatemplate
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {

	// Adding data to datatemplate
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello world"

	// Pulling data from session in appConfig and storing in datatenplate
	remoteIP := repo.AppConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// render template & inject datatemplate
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
