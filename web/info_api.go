package web

import (
	"context"

	"github.com/lgirma/gofx/common"
	"github.com/lgirma/gofx/config"
)

func InfoApi(baseUrl string, errorHandler ErrorHandler, webServer WebServer, infoService common.InfoService) {
	webServer.GET(baseUrl+"/ping", func(ctx context.Context) {
		apiKey := webServer.GetQuery(ctx, "apiKey")
		if apiKey != "f47ad" {
			webServer.RespondString(ctx, 400, "invalid api key")
			return
		}
		info := infoService.GetAppInfo()
		if info.Env == config.EnvProduction {
			webServer.RespondJson(ctx, 200, map[string]any{
				"version": info.Version,
			})
		} else {
			webServer.RespondJson(ctx, 200, map[string]any{
				"version": info.Version,
				"env":     info.Env,
			})
		}
	})

}
