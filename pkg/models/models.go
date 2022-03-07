package models

type Reservation struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	CheckIn string `json:"check_in"`
	CheckOut string `json:"check_out"`
	Adults string `json:"adults"`
	Children string `json:"children"`
	Room string `json:"room"`
}

