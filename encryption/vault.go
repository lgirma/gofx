package encryption

import (
	"github.com/lgirma/gofx/common"
	"os"
	"path/filepath"
)

type Vault interface {
	Store(content string, password string) error
	Read(password string) (string, error)
}

type FileVaultOptions struct {
	IsForCurrentUserOnly bool
	FileName             string
}

type FileVault struct {
	filePath string
}

func (f *FileVault) Store(content string, password string) error {
	enc, err := EncryptString(content, password)
	if err != nil {
		return err
	}
	err = os.WriteFile(f.filePath, []byte(enc), 0600)
	return err
}

func (f *FileVault) Read(password string) (string, error) {
	enc, err := os.ReadFile(f.filePath)
	if err != nil {
		return "", err
	}
	dec, err := DecryptString(string(enc), password)
	return dec, err
}

func NewFileVault(options FileVaultOptions, appInfo common.InfoService, dataDirs common.DataDirsService) (Vault, error) {
	dir := ""
	var err error
	if options.IsForCurrentUserOnly {
		dir, err = dataDirs.GetUserAppDataDir(appInfo.GetAppInfo().Name, false)
		if err != nil {
			return nil, err
		}
	} else {
		dir, err = dataDirs.GetSystemAppDataDir(appInfo.GetAppInfo().Name, false)
		if err != nil {
			return nil, err
		}
	}
	if options.FileName == "" {
		options.FileName = "lk.txt"
	}
	return &FileVault{
		filePath: filepath.Join(dir, options.FileName),
	}, nil
}
