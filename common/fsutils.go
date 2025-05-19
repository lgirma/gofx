package common

import (
	"os"
	"path/filepath"
	"strings"
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

func GetFileNameWithoutExtension(filePath string) string {
	fileNameWithExt := filepath.Base(filePath)
	fileExt := filepath.Ext(fileNameWithExt)
	fileNameWithoutExt := strings.TrimSuffix(fileNameWithExt, fileExt)
	return fileNameWithoutExt
}
