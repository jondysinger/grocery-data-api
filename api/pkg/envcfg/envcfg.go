package envcfg

import (
	"log"
	"os"
)

type EnvCfg struct {
	Port                  string
	KrogerApiBaseUrl      string
	KrogerApiClientId     string
	KrogerApiClientSecret string
	KrogerApiChain        string
	GroceryDataAppUrl     string
}

func Get() *EnvCfg {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("warning: %s environment variable not set", k)
		}
		return v
	}

	var cfg EnvCfg

	// Get environment variables
	cfg.Port = mustGetenv("PORT")
	cfg.KrogerApiBaseUrl = mustGetenv("KROGER_API_BASE_URL")
	cfg.KrogerApiClientId = mustGetenv("KROGER_API_CLIENT_ID")
	cfg.KrogerApiClientSecret = mustGetenv("KROGER_API_CLIENT_SECRET")
	cfg.KrogerApiChain = mustGetenv("KROGER_API_CHAIN")
	cfg.GroceryDataAppUrl = mustGetenv("GROCERY_DATA_APP_URL")

	return &cfg
}
