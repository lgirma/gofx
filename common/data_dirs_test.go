package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRootDataDirs(t *testing.T) {
	service := NewDataDirsService()
	dir, err := service.GetSystemAppDataRootDir()
	assert.NoError(t, err)
	assert.NotEmpty(t, dir)
	dir, err = service.GetUserAppDataRootDir()
	assert.NoError(t, err)
	assert.NotEmpty(t, dir)
}

func TestGetUserAppDataDir(t *testing.T) {
	service := NewDataDirsService()
	dir, err := service.GetUserAppDataDir("myTestApp", false)
	assert.Nil(t, err)
	assert.NotEmpty(t, dir)
	removeErr := service.RemoveUserAppDataDir("myTestApp")
	assert.Nil(t, removeErr)
}

func TestGetSystemAppDataDir(t *testing.T) {
	service := NewDataDirsService()
	dir, err := service.GetSystemAppDataDir("myTestApp", false)
	assert.Nil(t, err)
	assert.NotEmpty(t, dir)
	removeErr := service.RemoveSystemAppDataDir("myTestApp")
	assert.Nil(t, removeErr)
}
