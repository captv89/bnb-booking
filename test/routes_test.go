package test

import (
	"testing"

	"github.com/captv89/bnb-booking/pkg/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		t.Log("Handler is chi.Mux")
	default:
		t.Error("Handler is not chi.Mux", v)
	}
}
