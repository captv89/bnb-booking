package handler

import "github.com/captv89/bnb-booking/pkg/config"

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func AssignRepo (r *Repository) {
	Repo = r
}