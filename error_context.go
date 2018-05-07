import (
	"fmt"
	"os"
)

// START OMIT
func readConfig(configFile string) error {
	f, err := os.Open(configFile)
	if err != nil {
		return fmt.Errorf("could not read config from file %s: %v", configFile, err)
	}
	// do something with the open *File f
}

// END OMIT
