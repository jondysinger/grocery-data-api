package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jondysinger/grocery-data/api/cmd/api"
	"github.com/jondysinger/grocery-data/api/pkg/envcfg"
)

func main() {
	var app api.App

	// Get environment variables
	app.Config = envcfg.Get()

	// Start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%s", app.Config.Port), app.Routes())
	if err != nil {
		log.Fatal(err)
	}
}
