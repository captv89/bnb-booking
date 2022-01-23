package handler

import (
	"log"
	"net/http"

	"github.com/captv89/bnb-booking/pkg/models"
	"github.com/captv89/bnb-booking/pkg/render"
)

func (m *Repository) Book(w http.ResponseWriter, r *http.Request) {
	log.Println("Booking Page")

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap := make(map[string]string)
	stringMap["title"] = "About"
	stringMap["body"] = "This is the about page :)"
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "booking.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
