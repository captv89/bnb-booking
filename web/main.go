package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/captv89/bnb-booking/pkg/config"
	"github.com/captv89/bnb-booking/pkg/handler"
	"github.com/captv89/bnb-booking/pkg/models"
	"github.com/captv89/bnb-booking/pkg/render"
)

// Mention port number without localhost while running in production
var portNumber = "localhost:8080"

// Assign configuration struct
var app config.AppConfig

var session *scs.SessionManager

func main() {
	// Run the application
	err := run() 
	if err != nil {
		log.Fatal(err)
	}
	
	// Server routes
	log.Println("Server starting on port", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	// Start the server
	log.Fatal(srv.ListenAndServe())
}

func run() error {
	log.Println("Starting Application..")

	// what to put in the sessions to store and retrive data
	gob.Register(models.Reservation{})

	app.IsProduction = false

	// Session manager
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction
	app.Session = session

	// Load template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		return err
	}

	// Assign settings and values to config
	app.TemplateCache = tc
	app.UseCache = false

	// Pass config to handlers repo
	repo := handler.NewRepository(&app)
	handler.AssignRepo(repo)

	// Pass the cache to the render package
	render.GetTemplateCache(&app)

	return nil
}