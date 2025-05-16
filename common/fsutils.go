package common

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetAppDataDir() string {
	if runtime.GOOS == "windows" {
		return "%PROGRAMDATA%"
	} else if runtime.GOOS == "darwin" {
		return ""
	}
	return "/usr/share"
}

func InitMyAppDataDir(appName string, createIfNotExists bool) (string, error) {
	rootDir := GetAppDataDir()
	dir := filepath.Join(rootDir, appName)
	if createIfNotExists {
		if !FileExists(dir) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return "", err
			}
		}
	}
	return dir, nil
}

func GetMyAppDataDir(appName string) string {
	return filepath.Join(GetAppDataDir(), appName)
}

func GetUserAppDataDir() (string, error) {
	if runtime.GOOS == "windows" {
		return "%APPDATA%", nil
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".local/share"), nil
}

func InitMyAppUserDataDir(appName string, createIfNotExists bool) (string, error) {
	rootDir, err := GetUserAppDataDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(rootDir, appName)
	if createIfNotExists {
		if !FileExists(dir) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return "", err
			}
		}
	}
	return dir, nil
}

func GetMyAppUserDataDir(appName string) string {
	dir, _ := GetUserAppDataDir()
	return filepath.Join(dir, appName)
}

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
