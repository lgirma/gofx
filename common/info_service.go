package common

import (
	"github.com/lgirma/gofx/config"
)

type AppInfo struct {
	Version  string
	Env      string
	Name     string
	FullName string
	Edition  string
}

type InfoService interface {
	GetAppInfo() *AppInfo
}

type InfoServiceImpl struct {
	info AppInfo
}

func (i *InfoServiceImpl) GetAppInfo() *AppInfo {
	return &i.info
}

func NewAppInfo(appConfig config.AppConfig, appInfo AppInfo) InfoService {
	appInfo.Env = appConfig.GetEnv()
	return &InfoServiceImpl{
		info: appInfo,
	}
}
