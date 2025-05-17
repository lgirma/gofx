package common

import (
	"os"
	"path/filepath"
	"runtime"
)

type DataDirsService interface {
	GetSystemAppDataRootDir() string
	GetUserAppDataRootDir() string
	GetSystemAppDataDir(appName string, dontCreateIfNotExits bool) (string, error)
	GetUserAppDataDir(appName string, dontCreateIfNotExists bool) (string, error)
	RemoveSystemAppDataDir(appName string) error
	RemoveUserAppDataDir(appName string) error
}

type DefaultDataDirsService struct {
}

func (d *DefaultDataDirsService) RemoveSystemAppDataDir(appName string) error {
	dir, err := d.GetSystemAppDataDir(appName, true)
	if err != nil {
		return err
	}
	return os.RemoveAll(dir)
}

func (d *DefaultDataDirsService) RemoveUserAppDataDir(appName string) error {
	dir, err := d.GetUserAppDataDir(appName, true)
	if err != nil {
		return err
	}
	return os.RemoveAll(dir)
}

func (d *DefaultDataDirsService) GetSystemAppDataRootDir() string {
	if runtime.GOOS == "windows" {
		return "%PROGRAMDATA%"
	} else if runtime.GOOS == "darwin" {
		return "/Library/Application Support/"
	}
	return "/usr/share"
}

func (d *DefaultDataDirsService) GetUserAppDataRootDir() string {
	if runtime.GOOS == "windows" {
		return "%APPDATA%"
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "~"
	}
	return filepath.Join(homeDir, ".local/share")
}

func (d *DefaultDataDirsService) GetSystemAppDataDir(appName string, dontCreateIfNotExists bool) (string, error) {
	rootDir := d.GetSystemAppDataRootDir()
	dir := filepath.Join(rootDir, appName)
	if !dontCreateIfNotExists {
		if !FileExists(dir) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return "", err
			}
		}
	}
	return dir, nil
}

func (d *DefaultDataDirsService) GetUserAppDataDir(appName string, dontCreateIfNotExists bool) (string, error) {
	rootDir := d.GetUserAppDataRootDir()
	dir := filepath.Join(rootDir, appName)
	if !dontCreateIfNotExists {
		if !FileExists(dir) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return "", err
			}
		}
	}
	return dir, nil
}

func NewDataDirsService() DataDirsService {
	return &DefaultDataDirsService{}
}
