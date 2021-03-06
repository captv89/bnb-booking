package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/captv89/bnb-booking/pkg/config"
	"github.com/captv89/bnb-booking/pkg/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func GetTemplateCache(a *config.AppConfig) {
	app = a
}

//  AddDefaultData: Adds data for all templates
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Token = nosurf.Token(r)
	return td
}

// RenderTemplate: renders a template with the given data.
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data *models.TemplateData) {
	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Template not found: ", tmpl)
	}

	// log.Println("Debug t, ok, tmpl:", t, ok ,tmpl)

	buf := new(bytes.Buffer)

	data = AddDefaultData(data, r)

	err = t.Execute(buf, data)
	if err != nil {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)

	// Look for all page templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	// Add pages to cache
	for _, page := range pages {
		name := filepath.Base(page)
		log.Println("Adding template:", name)
		tmpl := template.Must(template.New(name).Funcs(functions).ParseFiles(page))

		// Look for all layout templates
		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// Add layouts to cache
		// log.Println("Adding layout templates:", layouts)

		if len(layouts) > 0 {
			tmpl = template.Must(tmpl.ParseGlob("./templates/*.layout.tmpl"))
		}

		// Adding the layout to the page
		myCache[name] = tmpl
	}
	return myCache, nil
}
