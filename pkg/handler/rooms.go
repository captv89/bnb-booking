package handler

import (
	"log"
	"net/http"

	"github.com/captv89/bnb-booking/pkg/models"
	"github.com/captv89/bnb-booking/pkg/render"
)

func (m *Repository) Armani(w http.ResponseWriter, r *http.Request) {
	log.Println("Room Armani Page")

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap := make(map[string]string)
	stringMap["title"] = "About"
	stringMap["body"] = "This is the about page :)"
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "armani.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Gucci(w http.ResponseWriter, r *http.Request) {
	log.Println("Room Gucci Page")

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap := make(map[string]string)
	stringMap["title"] = "About"
	stringMap["body"] = "This is the about page :)"
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "gucci.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Malabar(w http.ResponseWriter, r *http.Request) {
	log.Println("Cottage Malabar Page")

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap := make(map[string]string)
	stringMap["title"] = "About"
	stringMap["body"] = "This is the about page :)"
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "malabar.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}