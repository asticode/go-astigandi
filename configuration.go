package astigandi

import (
	"flag"

	astihttp "github.com/asticode/go-astitools/http"
)

// Flags
var (
	APIKey = flag.String("gandi-api-key", "", "the API key")
)

// Configuration represents the lib's configuration
type Configuration struct {
	APIKey string `toml:"api_key"`
	Sender astihttp.SenderOptions
}

// FlagConfig generates a Configuration based on flags
func FlagConfig() Configuration {
	return Configuration{
		APIKey: *APIKey,
	}
}
