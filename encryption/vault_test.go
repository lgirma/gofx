package encryption

import (
	"github.com/lgirma/gofx/common"
	"github.com/lgirma/gofx/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileVault(t *testing.T) {
	password := "myLongPassword"
	prize := "storedInVault"
	opts := FileVaultOptions{
		IsForCurrentUserOnly: true,
	}
	appInfo := common.NewAppInfo(config.NewStaticConfig(map[string]any{}), common.AppInfo{
		Version:  "",
		Env:      "",
		Name:     "myTestApp",
		FullName: "",
		Edition:  "",
	})
	dataDirs := common.NewDataDirsService()
	vault, err := NewFileVault(opts, appInfo, dataDirs)
	assert.Nil(t, err)
	err = vault.Store(prize, password)
	assert.NoError(t, err)

	dec, err := vault.Read(password)
	assert.NoError(t, err)
	assert.Equal(t, prize, dec)
}
