package main

import (
	"net/http"
	"strconv"

	"github.com/jondysinger/grocery-data-api/pkg/kclient"
)

// Creates a Kroger API client and retrieves an OAuth2 token
func (app *application) setupClient() (*kclient.KClient, error) {
	var client *kclient.KClient
	var err error

	// Create a client for the Kroger API
	client, err = kclient.New(
		app.Config.KrogerApiBaseUrl,
		app.Config.KrogerApiClientId,
		app.Config.KrogerApiClientSecret,
		app.Config.KrogerApiChain,
	)
	if err != nil {
		return nil, err
	}

	// Authorize the client
	if err := client.GetAuthToken(); err != nil {
		return nil, err
	}

	return client, nil
}

// Gets locations based on zip
func (app *application) Locations(w http.ResponseWriter, r *http.Request) {
	zipcode := r.URL.Query().Get("zipcode")
	filterLimit := r.URL.Query().Get("filterLimit")

	filterLimitConv, err := strconv.Atoi(filterLimit)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	// Setup the KClient
	client, err := app.setupClient()
	if err != nil {
		app.errorJson(w, err, http.StatusInternalServerError)
		return
	}

	// Get locations based on zip
	locations, err := client.GetLocations(zipcode, filterLimitConv)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	// Write the json response
	_ = app.writeJson(w, http.StatusOK, locations)
}

// Gets products based on filter and location
func (app *application) Products(w http.ResponseWriter, r *http.Request) {
	filterTerm := r.URL.Query().Get("filterTerm")
	locationId := r.URL.Query().Get("locationId")
	filterOffset := r.URL.Query().Get("filterOffset")
	filterLimit := r.URL.Query().Get("filterLimit")

	filterOffsetConv, err := strconv.Atoi(filterOffset)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	filterLimitConv, err := strconv.Atoi(filterLimit)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	// Setup the KClient
	client, err := app.setupClient()
	if err != nil {
		app.errorJson(w, err, http.StatusInternalServerError)
		return
	}

	// Get a list of products by filter and location
	products, err := client.GetProducts(filterTerm, locationId, filterOffsetConv, filterLimitConv)
	if err != nil {
		app.errorJson(w, err, http.StatusBadRequest)
		return
	}

	// Write the json response
	_ = app.writeJson(w, http.StatusOK, products)
}
