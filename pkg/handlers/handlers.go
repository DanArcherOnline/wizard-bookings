package handlers

import (
	"net/http"

	"github.com/danarcheronline/wizard-bookings/pkg/models"

	"github.com/danarcheronline/wizard-bookings/pkg/config"

	"github.com/danarcheronline/wizard-bookings/pkg/render"
)

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//Repo the repository used by handlers
var Repo *Repository

//NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "I am a testingmathingymabobything!"

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

//About is the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
