package dbrepo

import (
	"database/sql"

	"github.com/captv89/bnb-booking/pkg/config"
	"github.com/captv89/bnb-booking/pkg/repo"
)



type postgresDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewPostgresRepo (conn *sql.DB, app *config.AppConfig) repo.DatabaseRepo {
	return &postgresDBRepo{
		App: app,
		DB: conn,
	}
}