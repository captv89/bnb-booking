package handler

import (
	"log"
	"net/http"

	"github.com/captv89/bnb-booking/pkg/models"
	"github.com/captv89/bnb-booking/pkg/render"
)

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home Page")

	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	stringMap := make(map[string]string)
	stringMap["title"] = "Home"
	stringMap["body"] = "This is the home page :)"
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
