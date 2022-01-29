package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/captv89/bnb-booking/pkg/forms"
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
	// log.Println("Booking Page Post")
	// checkin := r.FormValue("checkin")
	// checkout := r.FormValue("checkout")
	// _, _ = w.Write([]byte(fmt.Sprintf("Start Date is %s and End Date is %s", checkin, checkout)))
	
	// Parse the form
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	adults, err := strconv.ParseInt(r.PostFormValue("adults"), 0, 0)
	children, err := strconv.ParseInt(r.PostFormValue("children"), 0, 0)

	reservation := models.Reservation{
		Room: r.FormValue("room_select"),
		FirstName: r.FormValue("first_name"),
		LastName: r.FormValue("last_name"),
		Email: r.FormValue("email"),
		Phone: r.FormValue("phone"),
		CheckIn: r.FormValue("check_in"),
		CheckOut: r.FormValue("check_out"),
		Adults: int(adults),
		Children: int(children),
	}

	form := forms.New(r.PostForm)

	form.Has("first_name", r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "booking.page.tmpl", &models.TemplateData{
		Form: form,
		Data: data,
	})
	return
	}
	
	
	
}
