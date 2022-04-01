package dbrepo

import (
	"context"
	"time"

	"github.com/captv89/bnb-booking/pkg/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

//Inserts a reservation into database
func (m *postgresDBRepo) InsertReservation(res *models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stmt := `INSERT INTO reservations (user_id, adults, children, check_in, check_out, room_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	
	_, err := m.DB.ExecContext(ctx, stmt, res.UserID, res.Adults, res.Children, res.CheckIn, res.CheckOut, res.RoomID, res.CreatedAt, res.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}