package render

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"github.com/ZackDiego/bookings/pkg/config"
	"github.com/ZackDiego/bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the application config
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// Check if app is nil before accessing TemplateCache
	if app == nil {
		http.Error(w, "app config not initialized", http.StatusInternalServerError)
		return
	}
     
     var tc map[string]*template.Template
	// create template cache
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		http.Error(w, "Error writing template", http.StatusInternalServerError)
		return
	}
}





func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	parsedTemplate, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range parsedTemplate {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		match, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(match) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
