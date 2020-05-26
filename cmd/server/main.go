package main

import (
	"fmt"
	"go-pg-domain-hex-actor/pkg/adding"
	"go-pg-domain-hex-actor/pkg/http/rest"
	"go-pg-domain-hex-actor/pkg/listing"
	"go-pg-domain-hex-actor/pkg/storage"
	"log"
	"net/http"
)

func main() {
	// In this repository we use Postgres
	var adder adding.Service
	var lister listing.Service

	// We will hardcode your connection here.
	// In your project we can create config and get it from it.
	// Database URL: host=localhost port=5432 user=poastgres password=your_password dbname=go_pg_domain_hex_actor sslmode=disable
	postgresURL := "host=localhost port=5432 user=postgres password=your_password dbname=go_pg_domain_hex_actor sslmode=disable"
	s, err := storage.NewStorage(postgresURL)
	if err != nil {
		panic(err)
	}

	adder = adding.NewService(s)
	lister = listing.NewService(s)

	// set up the HTTP server
	router := rest.Handler(adder, lister)

	fmt.Println("The beer server is on tap now: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
