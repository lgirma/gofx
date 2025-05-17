package common

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

const ErrDataDirRootNotFound = "DataDirRootNotFoundError"

type DataDirsService interface {
	GetSystemAppDataRootDir() (string, error)
	GetUserAppDataRootDir() (string, error)
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

func (d *DefaultDataDirsService) GetSystemAppDataRootDir() (string, error) {
	result := ""
	if runtime.GOOS == "windows" {
		result = os.ExpandEnv("$PROGRAMDATA")
	} else if runtime.GOOS == "darwin" {
		result = "/Library/Application Support/"
	} else {
		result = "/usr/share"
	}
	if !FileExists(result) {
		return "", errors.New(ErrDataDirRootNotFound)
	}
	return result, nil
}

func (d *DefaultDataDirsService) GetUserAppDataRootDir() (string, error) {
	result := ""
	if runtime.GOOS == "windows" {
		result = os.ExpandEnv("$APPDATA")
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			homeDir = os.ExpandEnv("$HOME")
		}
		result = filepath.Join(homeDir, ".local/share")
	}
	
	if !FileExists(result) {
		return "", errors.New(ErrDataDirRootNotFound)
	}
	return result, nil
}

func (d *DefaultDataDirsService) GetSystemAppDataDir(appName string, dontCreateIfNotExists bool) (string, error) {
	rootDir, err := d.GetSystemAppDataRootDir()
	if err != nil {
		return "", err
	}
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
	rootDir, err := d.GetUserAppDataRootDir()
	if err != nil {
		return "", err
	}
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
