package envcfg

import (
	"fmt"
	"os"
)

type EnvCfg struct {
	BaseUrl      string
	ClientId     string
	ClientSecret string
	Chain        string
}

func Get() *EnvCfg {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			panic(fmt.Sprintf("warning: %s environment variable not set", k))
		}
		return v
	}

	var cfg EnvCfg

	// Get environment variables
	cfg.BaseUrl = mustGetenv("KROGER_API_BASE_URL")
	cfg.ClientId = mustGetenv("KROGER_API_CLIENT_ID")
	cfg.ClientSecret = mustGetenv("KROGER_API_CLIENT_SECRET")
	cfg.Chain = mustGetenv("KROGER_API_CHAIN_FRED")

	return &cfg
}
