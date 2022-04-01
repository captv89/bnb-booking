package models

import "time"

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	AccessLevel string `json:"access_level"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Room struct {
	ID int `json:"id"`
	RoomName string `json:"room_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Restriction struct {
	ID int `json:"id"`
	RestrictionName string `json:"restriction_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Reservation struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Adults int `json:"adults"`
	Children int `json:"children"`
	CheckIn time.Time `json:"check_in"`
	CheckOut time.Time `json:"check_out"`
	RoomID int `json:"room_id"`
	User User `json:"user"`
	Room Room `json:"room"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoomRestriction struct {
	ID int `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	RoomID int `json:"room_id"`
	ReservationID int `json:"reservation_id"`
	RestrictionID int `json:"restriction_id"`
	Room Room `json:"room"`
	Reservation Reservation `json:"reservation"`
	Restriction Restriction `json:"restriction"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}