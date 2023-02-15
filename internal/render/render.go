package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/paulbyrne/bookingapp/internal/config"
	"github.com/paulbyrne/bookingapp/internal/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the templates packages
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(tempData *models.TemplateData, r *http.Request) *models.TemplateData {
	tempData.CSRFToken = nosurf.Token(r)
	return tempData
}

// RenderTemplate parses the html template and renders it
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, tempData *models.TemplateData) {

	var tempCache map[string]*template.Template

	if app.UseCache {
		// get template cache from app config
		tempCache = app.TemplateCache
	} else {
		tempCache, _ = CreateTemplateCache()
	}

	// retrieve template from cache
	template, inCache := tempCache[tmpl]
	if !inCache {
		log.Fatal("could not get template from template cache")
	}

	// create buffer
	buf := new(bytes.Buffer)

	tempData = AddDefaultData(tempData, r)

	err := template.Execute(buf, tempData)
	if err != nil {
		log.Println(err)
	}

	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	// create empty map - key (string) : value (pointer to template)
	tempCache := map[string]*template.Template{}

	// get all files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("../../templates/*.page.tmpl")
	if err != nil {
		return tempCache, err
	}

	// iterate through files ending with *.page.tmpl
	for _, page := range pages {

		name := filepath.Base(page)                    // strip path - retrieve name only
		ts, err := template.New(name).ParseFiles(page) // ts - template set
		if err != nil {
			return tempCache, err
		}

		// look for layout files
		matches, err := filepath.Glob("../../templates/*.layout.tmpl")
		if err != nil {
			return tempCache, err
		}

		// if layout files found - parse them
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.tmpl")
			if err != nil {
				return tempCache, err
			}
		}

		// Add template set (pointer to template)
		tempCache[name] = ts
	}

	return tempCache, nil
}
