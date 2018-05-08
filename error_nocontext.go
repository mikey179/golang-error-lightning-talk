import (
	"os"
)

// START OMIT
func readConfig(configFile string) (*Config, error) {
	f, err := os.Open(configFile)
	if err != nil {
		return nil err
	}
	// do something with the open *File f and return Config
}

// END OMIT
