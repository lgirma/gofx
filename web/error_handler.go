package web

import (
	"context"
	"errors"
	"log/slog"

	"github.com/lgirma/gofx/common"
)

type ErrorHandler interface {
	HandleError(err error, ctx context.Context)
	HandleIgnorableError(err error, ctx context.Context)
	HandleInputError(err error, ctx context.Context)
}

func NewErrorHandler(logger *slog.Logger, webServer WebServer) ErrorHandler {
	return &DefaultErrorHandler{
		Logger: logger,
		server: webServer,
	}
}

type DefaultErrorHandler struct {
	Logger *slog.Logger
	server WebServer
}

func (handler *DefaultErrorHandler) HandleIgnorableError(err error, c context.Context) {
	if err == nil {
		return
	}
	handler.Logger.Error("ignorable error", "reason", err)
	handler.server.RespondNoContent(c)
}

func (handler *DefaultErrorHandler) HandleInputError(err error, c context.Context) {
	if err == nil {
		return
	}
	handler.Logger.Error("input error", "reason", err)
	handler.server.RespondString(c, 400, "")
}

func (handler *DefaultErrorHandler) HandleError(err error, c context.Context) {
	if err == nil {
		return
	}
	var userError *common.UserError
	if errors.As(err, &userError) {
		handler.Logger.Error("user error",
			"code", userError.Code,
			"details", userError.Detail,
			"reason", userError.InternalError,
		)
		handler.server.RespondJson(c, 400, map[string]any{
			"Code":    userError.Code,
			"Details": userError.Detail,
		})
	} else {
		handler.Logger.Error("internal server error", "reason", err)
		handler.server.RespondJson(c, 500, map[string]any{
			"Code": common.ErrInternalServerError,
		})
	}
}
