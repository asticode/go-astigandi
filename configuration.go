package astigandi

import (
	"flag"

	"github.com/asticode/go-astikit"
)

// Flags
var (
	APIKey = flag.String("gandi-api-key", "", "the API key")
)

// Configuration represents the lib's configuration
type Configuration struct {
	APIKey string `toml:"api_key"`
	Sender astikit.HTTPSenderOptions
}

// FlagConfig generates a Configuration based on flags
func FlagConfig() Configuration {
	return Configuration{
		APIKey: *APIKey,
	}
}
