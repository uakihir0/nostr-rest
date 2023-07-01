package server

import (
	"github.com/labstack/echo/v4"
	"github.com/uakihir0/nostr-rest/server/api"
	"github.com/uakihir0/nostr-rest/server/mastodon/api"
	"github.com/uakihir0/nostr-rest/server/mastodon/openapi"
	"github.com/uakihir0/nostr-rest/server/openapi"
)

type Handler struct {
	api.SimpleHandler
	mapi.MastodonHandler
}

func RegisterHandlers(e *echo.Echo, h Handler) {
	openapi.RegisterHandlers(e, &h.SimpleHandler)
	mopenapi.RegisterHandlers(e, &h.MastodonHandler)
	e.GET("/health", GetHealth)
}
