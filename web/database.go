package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func database() {

	// structure the db uri
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_HOST"),os.Getenv("DB_PORT"),os.Getenv("DB_NAME"))

	//connect to database
	conn, err := pgx.Connect(context.Background(), dbURI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	log.Println("Connected to database")

	// test the new connection
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Cannot ping the DB", err)
	} else {
		log.Println("Successfully pinged DB")
	}
}

