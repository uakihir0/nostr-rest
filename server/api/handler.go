package api

import (
	"github.com/labstack/echo/v4"
	"github.com/uakihir0/nostr-rest/server/openapi"
)

type Handler struct {
}

func (h *Handler) RegisterHandler(e *echo.Echo) {
	openapi.RegisterHandlers(e, h)
}
