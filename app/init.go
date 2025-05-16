package app

import "github.com/lgirma/gofx/common"

type InitOptions struct {
	UseUserAppDataDir   bool
	UseSystemAppDataDir bool
}

type App interface {
	Init(options InitOptions) error
	GetUserAppDataDir() string
	GetSystemAppDataDir() string
}

type DefaultInitializer struct {
	appInfo          common.InfoService
	UserAppDataDir   string
	SystemAppDataDir string
}

func (i *DefaultInitializer) GetUserAppDataDir() string {
	return i.UserAppDataDir
}

func (i *DefaultInitializer) GetSystemAppDataDir() string {
	return i.SystemAppDataDir
}

func (i *DefaultInitializer) Init(options InitOptions) error {
	info := i.appInfo.GetAppInfo()
	if options.UseUserAppDataDir {
		userDataDir, err := common.InitMyAppUserDataDir(info.Name, true)
		if err != nil {
			return err
		}
		i.UserAppDataDir = userDataDir
	}
	if options.UseSystemAppDataDir {
		sysDataDir, err := common.InitMyAppDataDir(info.Name, true)
		if err != nil {
			return err
		}
		i.SystemAppDataDir = sysDataDir
	}
	return nil
}

func NewApp(appInfo common.InfoService) App {
	return &DefaultInitializer{
		appInfo: appInfo,
	}
}
