package handler

import (
	"log"
	"net/http"

	"github.com/captv89/bnb-booking/pkg/forms"
	"github.com/captv89/bnb-booking/pkg/models"
	"github.com/captv89/bnb-booking/pkg/render"
)

func (m *Repository) Book(w http.ResponseWriter, r *http.Request) {
	log.Println("Booking Page")
	data := make(map[string]interface{})
	emtyReservation := models.Reservation{}
	data["reservation"] = emtyReservation
	render.RenderTemplate(w, r, "booking.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repository) PostBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Booking Page Post")

	fields := []string{"first_name", "last_name", "email", "phone", "check_in", "check_out", "adults", "children"}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		Room:      r.Form.Get("room"),
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		CheckIn:   r.Form.Get("check_in"),
		CheckOut:  r.Form.Get("check_out"),
		Adults:    r.Form.Get("adults"),
		Children:  r.Form.Get("children"),
	}

	form := forms.New(r.PostForm)
	form.Validate(fields, r)
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, r, "booking.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/summary", http.StatusSeeOther)
}

func (m *Repository) BookingSummary(w http.ResponseWriter, r *http.Request) {
	log.Println("Booking Summary Page")
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.Session.Put(r.Context(), "error", "No reservation found")
		http.Redirect(w, r, "/book", http.StatusSeeOther)
		return
	}
	// If loading successful then remove the session
	m.App.Session.Remove(r.Context(), "reservation")
	
	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderTemplate(w, r, "summary.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}
