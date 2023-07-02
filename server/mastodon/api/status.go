package mapi

import (
	"github.com/labstack/echo/v4"
	"github.com/uakihir0/nostr-rest/server/mastodon/openapi"
)

func (h *MastodonHandler) GetApiV1AccountsUidStatuses(
	c echo.Context,
	uid string,
	params mopenapi.GetApiV1AccountsUidStatusesParams,
) error {
	//TODO implement me
	panic("implement me")
}
