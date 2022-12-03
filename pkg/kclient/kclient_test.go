package kclient

import (
	"os"
	"testing"

	"github.com/jondysinger/grocery-data-api/pkg/envcfg"
)

var cfg *envcfg.EnvCfg

func oneTimeSetup() {
	// Get environment variables
	cfg = envcfg.Get()
}

func oneTimeTeardown() {
}

func TestMain(m *testing.M) {
	oneTimeSetup()
	code := m.Run()
	oneTimeTeardown()
	os.Exit(code)
}

func TestNew(t *testing.T) {
	if _, err := New(cfg.BaseUrl, cfg.ClientId, cfg.ClientSecret, cfg.Chain); err != nil {
		t.Fatalf("expected success but got error, %v", err)
	}
}

func TestNewKClientInvalidParam(t *testing.T) {
	testCases := []struct {
		name   string
		url    string
		id     string
		secret string
		chn    string
	}{
		{"cfg.BaseUrl guard clause", "", cfg.ClientId, cfg.ClientSecret, cfg.Chain},
		{"id guard clause", cfg.BaseUrl, "", cfg.ClientSecret, cfg.Chain},
		{"secret guard clause", cfg.BaseUrl, cfg.ClientId, "", cfg.Chain},
		{"cfg.Chain guard clause", cfg.BaseUrl, cfg.ClientId, cfg.ClientSecret, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := New(tc.url, tc.id, tc.secret, tc.chn); err == nil {
				t.Error("expected error but was none")
			}
		})
	}
}

func TestGetAuthToken(t *testing.T) {
	client, err := New(cfg.BaseUrl, cfg.ClientId, cfg.ClientSecret, cfg.Chain)
	if err != nil {
		t.Fatalf("error during client setup, %v", err)
	}

	if err := client.GetAuthToken(); err != nil {
		t.Fatalf("expected success but got error, %v", err)
	}
}

func TestGetAuthTokenInvalidParam(t *testing.T) {
	testCases := []struct {
		name   string
		url    string
		id     string
		secret string
		chn    string
	}{
		{"url invalid", "badUrl", cfg.ClientId, cfg.ClientSecret, cfg.Chain},
		{"id invalid", cfg.BaseUrl, "12345678", cfg.ClientSecret, cfg.Chain},
		{"secret invalid", cfg.BaseUrl, cfg.ClientId, "12345678", cfg.Chain},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := New(tc.url, tc.id, tc.secret, tc.chn)
			if err != nil {
				t.Fatalf("failure in setup code, %v", err)
			}

			if err := client.GetAuthToken(); err == nil {
				t.Error("expected err but got none")
			}
		})
	}
}

func TestGetLocations(t *testing.T) {
	client, err := New(cfg.BaseUrl, cfg.ClientId, cfg.ClientSecret, cfg.Chain)
	if err != nil {
		t.Fatalf("error during client setup, %v", err)
	}
	if err := client.GetAuthToken(); err != nil {
		t.Fatalf("error during auth setup, %v", err)
	}

	locations, err := client.GetLocations("97224")
	if err != nil {
		t.Fatalf("expected success but got error, %v", err)
	} else if len(locations) < 1 {
		t.Fatalf("expected at least one location but got %d", len(locations))
	}
}

func TestGetLocationsInvalidParam(t *testing.T) {
	client, err := New(cfg.BaseUrl, cfg.ClientId, cfg.ClientSecret, cfg.Chain)
	if err != nil {
		t.Fatalf("error during client setup, %v", err)
	}
	if err := client.GetAuthToken(); err != nil {
		t.Fatalf("error during auth setup, %v", err)
	}
	testCases := []struct {
		name string
		zip  string
	}{
		{"zipCode too long", "123456"},
		{"zipCode too short", "1234"},
		{"zipCode not all digits", "T97224A"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := client.GetLocations(tc.zip); err == nil {
				t.Error("expected err but got none")
			}
		})
	}
}

func TestGetProducts(t *testing.T) {
	client, err := New(cfg.BaseUrl, cfg.ClientId, cfg.ClientSecret, cfg.Chain)
	if err != nil {
		t.Fatalf("error during client setup, %v", err)
	}
	if err := client.GetAuthToken(); err != nil {
		t.Fatalf("error during auth setup, %v", err)
	}

	products, err := client.GetProducts("milk", "", 5, 1)
	if err != nil {
		t.Fatalf("expected success but got error, %v", err)
	} else if len(products) != 1 {
		t.Fatalf("expected one product but got %d", len(products))
	}
}

func TestGetProductsInvalidParam(t *testing.T) {
	client, err := New(cfg.BaseUrl, cfg.ClientId, cfg.ClientSecret, cfg.Chain)
	if err != nil {
		t.Fatalf("error during client setup, %v", err)
	}
	if err := client.GetAuthToken(); err != nil {
		t.Fatalf("error during auth setup, %v", err)
	}
	testCases := []struct {
		name   string
		filter string
		loc    string
		offset int
		limit  int
	}{
		{"locationId invalid", "milk", "12345", 0, 1},
		{"offset invalid", "milk", "", -1, 1},
		{"limit invalid", "milk", "", 0, 50000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := client.GetProducts(tc.filter, tc.loc, tc.offset, tc.limit); err == nil {
				t.Error("expected err but got none")
			}
		})
	}
}
