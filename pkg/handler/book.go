package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/captv89/bnb-booking/pkg/forms"
	"github.com/captv89/bnb-booking/pkg/helpers"
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

	// Convert to Date format
	sd := r.Form.Get("check_in")
	ed := r.Form.Get("check_out")

	layout := "20.01.2006"
	checkIn, err := time.Parse(layout, sd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	checkOut, err := time.Parse(layout, ed)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// Covert String to int
	adults, err := strconv.Atoi(r.Form.Get("adults"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	children, err := strconv.Atoi(r.Form.Get("children"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	userID, err := strconv.Atoi(r.Form.Get("user_id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}


	// Create a new reservation object
	reservation := models.Reservation{
		CheckIn:   checkIn,
		CheckOut:  checkOut,
		Adults:    adults,
		Children:  children,
		UserID: userID,
		RoomID:  roomID,
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

	// Insert the reservation into database
	err = m.DB.InsertReservation(&reservation)
	if err != nil {
		helpers.ServerError(w, err)
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
