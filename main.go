package main

import (
	"log"
	"net/http"

	"github.com/ricardo-ch/after-sales-invoicing/src/api"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/config"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/database"
	"github.com/ricardo-ch/after-sales-invoicing/src/api/logger"

	"github.com/gorilla/mux"
)

func init() {
	//Get configuration
	config.Get()

	//Connect to databases
	database.Connect()
}

func main() {

	//defer the close of the databases
	defer database.Close()

	//Setup the API routes
	router := mux.NewRouter().StrictSlash(true)

	// Create Api routes
	api.NewRouter(router)

	log.Println("The comsab-app service has been started on port", *config.Port)

	log.Fatal(http.ListenAndServe(":"+*config.Port, router))
}
