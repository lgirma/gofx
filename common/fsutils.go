package common

import (
	"os"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	// For other errors, handle as needed (e.g., permission issues)
	return false
}
