package app

import "github.com/lgirma/gofx/common"

type InitOptions struct {
	UseUserAppDataDir   bool
	UseSystemAppDataDir bool
}

type App interface {
	Init(options InitOptions) error
}

type DefaultInitializer struct {
	appInfo common.InfoService
}

func (i *DefaultInitializer) Init(options InitOptions) error {
	return nil
}

func NewApp(appInfo common.InfoService) App {
	return &DefaultInitializer{
		appInfo: appInfo,
	}
}
