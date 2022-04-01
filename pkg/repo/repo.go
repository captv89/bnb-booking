package repo

import "github.com/captv89/bnb-booking/pkg/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res *models.Reservation) error
}