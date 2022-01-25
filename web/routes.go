package main

import (
	"net/http"

	"github.com/captv89/bnb-booking/pkg/config"
	"github.com/captv89/bnb-booking/pkg/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Middleware
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(LoadSession)

	// Routes
	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	mux.Get("/contact", handler.Repo.Contact)
	mux.Get("/book", handler.Repo.Book)
	mux.Post("/book", handler.Repo.PostBook)
	mux.Get("/room-armani", handler.Repo.Armani)
	mux.Get("/room-gucci", handler.Repo.Gucci)
	mux.Get("/cottage-malabar", handler.Repo.Malabar)
	mux.Post("/search-availability", handler.Repo.SearchAvailability)

	// Load static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
