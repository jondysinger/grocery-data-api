package main

import (
	"fmt"

	"github.com/jondysinger/grocery-data-api/pkg/envcfg"
	"github.com/jondysinger/grocery-data-api/pkg/kclient"
)

func main() {
	// Get environment variables
	cfg := envcfg.Get()

	// Create a client for the Kroger API
	client, err := kclient.New(cfg.BaseUrl, cfg.ClientId, cfg.ClientSecret, cfg.Chain)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize client, %v", err))
	}

	// Authorize the client
	if err := client.GetAuthToken(); err != nil {
		panic(fmt.Sprintf("failed to get auth token, %v", err))
	}

	// Get locations based on zip
	locations, err := client.GetLocations("97224")
	if err != nil {
		panic(fmt.Sprintf("failed to get locations, %v", err))
	}
	loc := locations[0]

	// Get a list of products by filter
	products, err := client.GetProducts("milk", loc.LocationId, 0, 50)
	if err != nil {
		panic(fmt.Sprintf("failed to get products, %v", err))
	}

	// List out the products to the terminal
	for _, product := range products {
		fmt.Printf("Product: '%s'\n", product.Description)
	}
}
