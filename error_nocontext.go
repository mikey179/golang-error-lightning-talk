import (
	"os"
)

// START OMIT
func readConfig(configFile string) error {
	f, err := os.Open(configFile)
	if err != nil {
		return err
	}
	// do something with the open *File f
}

// END OMIT
