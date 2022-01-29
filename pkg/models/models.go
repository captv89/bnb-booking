package models

type Reservation struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	CheckIn string `json:"check_in"`
	CheckOut string `json:"check_out"`
	Adults int `json:"adults"`
	Children int `json:"children"`
	Room string `json:"room"`
}