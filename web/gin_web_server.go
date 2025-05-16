package web

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lgirma/gofx/config"
)

type GinGonicWebServer struct {
	bindAddress string
	port        int64
	engine      *gin.Engine
}

func (server *GinGonicWebServer) StaticFS(rootUrl string, fs http.FileSystem) {
	server.engine.StaticFS(rootUrl, fs)
}

func (server *GinGonicWebServer) SetCorsOptions(options *CorsOptions) {
	corsConfig := cors.DefaultConfig()
	if options.AllowAllHeaders {
		corsConfig.AllowHeaders = []string{"*"}
	} else {
		corsConfig.AllowHeaders = options.AllowedHeaders
	}
	if options.AllowAllOrigins {
		corsConfig.AllowAllOrigins = true
	} else {
		corsConfig.AllowOrigins = options.AllowedOrigins
	}
	if options.AllowAllMethods {
		corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	} else {
		corsConfig.AllowMethods = options.AllowedMethods
	}
	server.engine.Use(cors.New(corsConfig))
}

func (server *GinGonicWebServer) Abort(ctx context.Context) {
	ctx.(*gin.Context).Abort()
}

func (server *GinGonicWebServer) BindBody(ctx context.Context, toObj any) error {
	return ctx.(*gin.Context).ShouldBind(toObj)
}

func (server *GinGonicWebServer) ConvertHandlers(handlers []func(ctx context.Context)) []gin.HandlerFunc {
	typedHandlers := []gin.HandlerFunc{}
	for i := range handlers {
		typedHandlers = append(typedHandlers, func(ctx *gin.Context) {
			handlers[i](ctx)
		})
	}
	return typedHandlers
}

func (server *GinGonicWebServer) DELETE(path string, handlers ...func(ctx context.Context)) {
	typedHandlers := server.ConvertHandlers(handlers)
	server.engine.DELETE(path, typedHandlers...)
}

func (server *GinGonicWebServer) FormFile(ctx context.Context, key string) (*multipart.FileHeader, error) {
	return ctx.(*gin.Context).FormFile(key)
}

func (server *GinGonicWebServer) GET(path string, handlers ...func(ctx context.Context)) {
	typedHandlers := server.ConvertHandlers(handlers)
	server.engine.GET(path, typedHandlers...)
}

func (server *GinGonicWebServer) GetHeader(ctx context.Context, key string) string {
	return ctx.(*gin.Context).GetHeader(key)
}

func (server *GinGonicWebServer) GetPathParam(ctx context.Context, key string) string {
	return ctx.(*gin.Context).Param(key)
}

func (server *GinGonicWebServer) GetPathParamInt64(ctx context.Context, key string) (int64, error) {
	str := ctx.(*gin.Context).Param(key)
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (server *GinGonicWebServer) GetSessionData(ctx context.Context, key string) (value any, exists bool) {
	return ctx.(*gin.Context).Get(key)
}

func (server *GinGonicWebServer) Next(ctx context.Context) {
	ctx.(*gin.Context).Next()
}

func (server *GinGonicWebServer) POST(path string, handlers ...func(ctx context.Context)) {
	typedHandlers := server.ConvertHandlers(handlers)
	server.engine.POST(path, typedHandlers...)
}

func (server *GinGonicWebServer) RespondBlob(ctx context.Context, code int, contentType string, fileName string, content []byte) {
	server.SetHeader(ctx, "Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	ctx.(*gin.Context).Data(code, contentType, content)
}

func (server *GinGonicWebServer) RespondJson(ctx context.Context, code int, content any) {
	ctx.(*gin.Context).JSON(code, content)
}

func (server *GinGonicWebServer) RespondNoContent(ctx context.Context) {
	ctx.(*gin.Context).Status(204)
}

func (server *GinGonicWebServer) RespondString(ctx context.Context, code int, content string) {
	ctx.(*gin.Context).String(code, content)
}

func (server *GinGonicWebServer) SaveUploadedFile(ctx context.Context, uploadedFile *multipart.FileHeader, destFilePath string) error {
	return ctx.(*gin.Context).SaveUploadedFile(uploadedFile, destFilePath)
}

func (server *GinGonicWebServer) SetHeader(ctx context.Context, key string, value string) {
	ctx.(*gin.Context).Header(key, value)
}

func (server *GinGonicWebServer) SetSessionData(ctx context.Context, key string, value any) {
	ctx.(*gin.Context).Set(key, value)
}

func (server *GinGonicWebServer) GetRequestInfo(ctx context.Context) *WebRequestInfo {
	ginCtx := ctx.(*gin.Context)
	return &WebRequestInfo{
		IpAddress: ginCtx.RemoteIP(),
		UserAgent: ginCtx.Request.UserAgent(),
	}
}

func (server *GinGonicWebServer) GetPort() int64 {
	return server.port
}

func (server *GinGonicWebServer) GetBindAddress() string {
	return server.bindAddress
}

func (server *GinGonicWebServer) SetPort(port int64) {
	server.port = port
}

func (server *GinGonicWebServer) SetBindAddress(bindAddress string) {
	server.bindAddress = bindAddress
}

func (server *GinGonicWebServer) Run() error {
	return server.engine.Run(fmt.Sprintf("%s:%d", server.bindAddress, server.port))
}

func (server *GinGonicWebServer) TryGetQuery(ctx context.Context, key string) (value string, exists bool) {
	return ctx.(*gin.Context).GetQuery(key)
}

func (server *GinGonicWebServer) GetQuery(ctx context.Context, key string) string {
	return ctx.(*gin.Context).Query(key)
}

func NewGinGonicWebServer(webServerConfig *WebServerConfig) WebServer {
	if webServerConfig.Env == config.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	}
	return &GinGonicWebServer{
		engine:      gin.Default(),
		bindAddress: webServerConfig.Bind,
		port:        webServerConfig.Port,
	}
}
