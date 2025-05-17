package licensing

import (
	"encoding/hex"
	"errors"
	"github.com/lgirma/gofx/encryption"
)

type CopyProtection interface {
	GetRequestCode() (string, error)
	GetActivationCode(requestCode string) (string, error)
	CheckActivationCode(requestCode string, activationCode string) (bool, error)
	IsLicensed() (bool, error)
	RemoveActivation() error
	ActivateLicense(activationCode string) error
}

type Options struct {
	AuthCode string
	Product  Product
}

type DefaultCopyProtection struct {
	options Options
	vault   encryption.Vault
}

func (d *DefaultCopyProtection) RemoveActivation() error {
	return d.vault.Store("", d.options.AuthCode)
}

func (d *DefaultCopyProtection) GetRequestCode() (string, error) {
	licenseId, err := GetLicenseId(d.options.Product, LicenseIdOptions{})
	if err != nil {
		return "", err
	}
	return hex.EncodeToString([]byte(licenseId)), nil
}

func (d *DefaultCopyProtection) GetActivationCode(requestCode string) (string, error) {
	return encryption.EncryptString(requestCode, d.options.AuthCode)
}

func (d *DefaultCopyProtection) CheckActivationCode(requestCode string, activationCode string) (bool, error) {
	dec, err := encryption.DecryptString(activationCode, d.options.AuthCode)
	if err != nil {
		return false, err
	}
	return requestCode == dec, nil
}

func (d *DefaultCopyProtection) IsLicensed() (bool, error) {
	if !d.vault.Exists() {
		return false, nil
	}
	dec, err := d.vault.Read(d.options.AuthCode)
	if err != nil {
		return false, err
	}
	reqCode, err := d.GetRequestCode()
	if err != nil {
		return false, err
	}
	return reqCode == dec, nil
}

func (d *DefaultCopyProtection) ActivateLicense(activationCode string) error {
	reqCode, err := d.GetRequestCode()
	if err != nil {
		return err
	}
	check, err := d.CheckActivationCode(reqCode, activationCode)
	if !check {
		return errors.New(ErrInvalidActivationCode)
	}
	return d.vault.Store(reqCode, d.options.AuthCode)
}

func NewCopyProtectionService(vault encryption.Vault, options Options) CopyProtection {
	return &DefaultCopyProtection{
		options: options,
		vault:   vault,
	}
}
