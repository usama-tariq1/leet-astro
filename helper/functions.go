package helper

import (
	"errors"
	"log"
	"os"
)

func Contains(pin string, heystack []string) bool {
	for _, b := range heystack {
		if b == pin {
			return true
		}
	}
	return false
}

func GetWD() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}

func FileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		// path/to/whatever exists
		return true

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		return false
	} else {
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		return true

	}
}
