package handler

import (
	"github.com/captv89/bnb-booking/pkg/config"
	"github.com/captv89/bnb-booking/pkg/driver"
	"github.com/captv89/bnb-booking/pkg/repo"
	"github.com/captv89/bnb-booking/pkg/repo/dbrepo"
)

type Repository struct {
	App *config.AppConfig
	DB repo.DatabaseRepo
}

var Repo *Repository

func NewRepository(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a, 
		DB: dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func AssignRepo (r *Repository) {
	Repo = r
}