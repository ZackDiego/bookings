package handlers

import (
	"net/http"
	"github.com/ZackDiego/bookings/pkg/render"
	"github.com/ZackDiego/bookings/pkg/config"
	"github.com/ZackDiego/bookings/pkg/models"
)



// The repository used by the handlers
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

func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)


	
	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, this is a test"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// pass data to template
	render.RenderTemplate(w, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}



