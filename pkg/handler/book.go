package handler

import (
	"fmt"
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
	render.RenderTemplate(w, r, "booking.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) PostBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Booking Page Post")
	start := r.FormValue("start")
	end := r.FormValue("end")

	_, _ = w.Write([]byte(fmt.Sprintf("Start Date is %s and End Date is %s", start, end)))
}
