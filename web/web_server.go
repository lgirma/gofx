package web

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/lgirma/gofx/config"
)

type WebRequestInfo struct {
	IpAddress string
	UserAgent string
}

type CorsOptions struct {
	AllowAllOrigins bool
	AllowAllHeaders bool
	AllowAllMethods bool
	AllowedOrigins  []string
	AllowedMethods  []string
	AllowedHeaders  []string
}

type WebRequestHandler = func(ctx context.Context)

type WebServerConfig struct {
	Env  config.Environment
	Bind string
	Port int64
	Cors *CorsOptions
}

func GetWebServerConfig(appConfig config.AppConfig) *WebServerConfig {
	result := config.GetConfigOrDefault(appConfig, "server", &WebServerConfig{
		Bind: "",
		Port: 5555,
	})
	result.Env = appConfig.GetEnv()
	return result
}

type WebServer interface {
	GetRequestInfo(ctx context.Context) *WebRequestInfo
	GetPort() int64
	GetBindAddress() string
	SetPort(port int64)
	SetBindAddress(bindAddress string)
	Run() error

	GET(path string, handlers ...WebRequestHandler)
	POST(path string, handlers ...WebRequestHandler)
	DELETE(path string, handlers ...WebRequestHandler)

	RespondString(ctx context.Context, code int, content string)
	RespondBlob(ctx context.Context, code int, contentType string, fileName string, content []byte)
	RespondJson(ctx context.Context, code int, content any)
	RespondNoContent(ctx context.Context)

	SetSessionData(ctx context.Context, key string, value any)
	GetSessionData(ctx context.Context, key string) (value any, exists bool)

	GetHeader(ctx context.Context, key string) string
	Abort(ctx context.Context)
	Next(ctx context.Context)
	SetHeader(ctx context.Context, key string, value string)

	BindBody(ctx context.Context, toObj any) error
	GetPathParam(ctx context.Context, key string) string
	GetPathParamInt64(ctx context.Context, key string) (int64, error)
	TryGetQuery(ctx context.Context, key string) (value string, exists bool)
	GetQuery(ctx context.Context, key string) string

	FormFile(ctx context.Context, key string) (*multipart.FileHeader, error)
	SaveUploadedFile(ctx context.Context, uploadedFile *multipart.FileHeader, destFilePath string) error
	SetCorsOptions(options *CorsOptions)

	StaticFS(rootUrl string, fs http.FileSystem)
}
