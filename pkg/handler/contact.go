package handler

import (
	"log"
	"net/http"

	"github.com/captv89/bnb-booking/pkg/models"
	"github.com/captv89/bnb-booking/pkg/render"
)

func (*Repository) Contact(w http.ResponseWriter, r *http.Request) {
	log.Println("Contact Page")
	stringMap := make(map[string]string)
	stringMap["title"] = "Contact"
	stringMap["body"] = "This is the contact page :)"
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
