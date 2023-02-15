package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/paulbyrne/bookingapp/internal/config"
	"github.com/paulbyrne/bookingapp/internal/models"
	"github.com/paulbyrne/bookingapp/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// AboutUs is the about page handler
func (repo *Repository) AboutUs(w http.ResponseWriter, r *http.Request) {

	// Adding data to datatemplate
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello world"

	// Pulling data from session in appConfig and storing in datatenplate
	remoteIP := repo.AppConfig.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// render template & inject datatemplate
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// SearchAvailability is the availability checker page handler
func (repo *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {

	// render template & inject datatemplate
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostCheckAvailability is the post check availability handler
func (repo *Repository) PostCheckAvailability(w http.ResponseWriter, r *http.Request) {

	start := r.Form.Get("start_date")
	end := r.Form.Get("end_date")

	w.Write([]byte(fmt.Sprintf("start date: %s, end date %s", start, end)))
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for room availability and returns json response
func (repo *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok:      true,
		Message: "1 Room available",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// MakeReservation is the reservation page handler
func (repo *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {

	// render template & inject datatemplate
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// ContactUs is the contact us page handler
func (repo *Repository) ContactUs(w http.ResponseWriter, r *http.Request) {

	// render template & inject datatemplate
	render.RenderTemplate(w, r, "contact-us.page.tmpl", &models.TemplateData{})
}

// BasicSuite is the basic suite page handler
func (repo *Repository) BasicSuite(w http.ResponseWriter, r *http.Request) {

	// render template & inject datatemplate
	render.RenderTemplate(w, r, "basic-suite.page.tmpl", &models.TemplateData{})
}

// KingsSuite is the kings suite page handler
func (repo *Repository) KingsSuite(w http.ResponseWriter, r *http.Request) {

	// render template & inject datatemplate
	render.RenderTemplate(w, r, "kings-suite.page.tmpl", &models.TemplateData{})
}
