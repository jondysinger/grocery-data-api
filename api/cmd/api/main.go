package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jondysinger/grocery-data-api/pkg/envcfg"
)

type application struct {
	Config *envcfg.EnvCfg
}

func main() {
	var app application

	// Get environment variables
	app.Config = envcfg.Get()

	// Start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%s", app.Config.Port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
