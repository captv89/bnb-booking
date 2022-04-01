package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/captv89/bnb-booking/pkg/config"
	"github.com/captv89/bnb-booking/pkg/driver"
	"github.com/captv89/bnb-booking/pkg/handler"
	"github.com/captv89/bnb-booking/pkg/models"
	"github.com/captv89/bnb-booking/pkg/render"
	"github.com/joho/godotenv"
)

// Mention port number without localhost while running in production
var portNumber = "localhost:8080"

// Assign configuration struct
var app config.AppConfig

var session *scs.SessionManager

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func main() {
	// Run the application
	db, err := run() 
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	
	// Server routes
	log.Println("Server starting on port", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	// Start the server
	log.Fatal(srv.ListenAndServe())
}

func run() (*driver.DB, error) {
	log.Println("Starting Application..")

	// what to put in the sessions to store and retrive data
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(models.RoomRestriction{})


	app.IsProduction = false

	// Session manager
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction
	app.Session = session

	// Connect to database
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_HOST"),os.Getenv("DB_PORT"),os.Getenv("DB_NAME"))
	db, err := driver.ConnectSQL(dbURI)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	

	// Load template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		return nil, err
	}

	// Assign settings and values to config
	app.TemplateCache = tc
	app.UseCache = false

	// Pass config to handlers repo
	repo := handler.NewRepository(&app, db)
	handler.AssignRepo(repo)

	// Pass the cache to the render package
	render.GetTemplateCache(&app)

	return db, nil
}