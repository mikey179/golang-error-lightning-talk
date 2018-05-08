import (
	"fmt"
	"os"
)

// START OMIT
func readConfig(configFile string) (*Config, error) {
	f, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("could not read config from file %s: %v", configFile, err)
	}
	// do something with the open *File f and return Config
}

// END OMIT
