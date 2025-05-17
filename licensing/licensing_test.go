package licensing

import (
	"github.com/lgirma/gofx/common"
	"github.com/lgirma/gofx/config"
	"github.com/lgirma/gofx/encryption"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestCopyProtectionService() (CopyProtection, error) {
	opts := encryption.FileVaultOptions{
		IsForCurrentUserOnly: true,
	}
	appInfo := common.NewAppInfo(config.NewStaticConfig(map[string]any{}), common.AppInfo{
		Version:  "",
		Env:      "",
		Name:     "MyApp",
		FullName: "",
		Edition:  "Standard",
	})
	dataDirs := common.NewDataDirsService()
	vault, err := encryption.NewFileVault(opts, appInfo, dataDirs)
	service := NewCopyProtectionService(vault, Options{
		Product:  Product{Name: "MyApp", Edition: "Standard"},
		AuthCode: "1234",
	})
	return service, err
}

func TestDefaultCopyProtection_GetRequestCode(t *testing.T) {
	service, err := getTestCopyProtectionService()
	assert.Nil(t, err)
	reqCode, err := service.GetRequestCode()
	assert.NoError(t, err)
	assert.NotEmpty(t, reqCode)
	t.Logf("ReqCode: %s", reqCode)
}

func TestDefaultCopyProtection_GetActivationCode(t *testing.T) {
	service, err := getTestCopyProtectionService()
	assert.Nil(t, err)
	reqCode, err := service.GetRequestCode()
	assert.NoError(t, err)
	assert.NotEmpty(t, reqCode)
	actCode, err := service.GetActivationCode(reqCode)
	assert.NoError(t, err)
	assert.NotEmpty(t, actCode)
	t.Logf("ActCode: %s", actCode)
}

func TestDefaultCopyProtection_CheckActivationCode(t *testing.T) {
	service, err := getTestCopyProtectionService()
	assert.Nil(t, err)
	reqCode, err := service.GetRequestCode()
	assert.NoError(t, err)
	actCode, err := service.GetActivationCode(reqCode)
	assert.NoError(t, err)
	checkResult, err := service.CheckActivationCode(reqCode, actCode)
	assert.NoError(t, err)
	assert.True(t, checkResult)
	checkResult, err = service.CheckActivationCode(reqCode, "wrongActivationCode")
	assert.False(t, checkResult)
}

func TestDefaultCopyProtection_RemoveActivation(t *testing.T) {
	service, err := getTestCopyProtectionService()
	assert.Nil(t, err)
	err = service.RemoveActivation()
	assert.Nil(t, err)
	res, err := service.IsLicensed()
	assert.NoError(t, err)
	assert.False(t, res)
}

func TestDefaultCopyProtection_IsLicensed(t *testing.T) {
	service, err := getTestCopyProtectionService()
	assert.NoError(t, err)
	_ = service.RemoveActivation()
	status, err := service.IsLicensed()
	assert.NoError(t, err)
	assert.False(t, status)

	reqCode, err := service.GetRequestCode()
	assert.NoError(t, err)
	actCode, err := service.GetActivationCode(reqCode)
	err = service.ActivateLicense(actCode)
	assert.NoError(t, err)

	status, err = service.IsLicensed()
	assert.NoError(t, err)
	assert.True(t, status)
}
