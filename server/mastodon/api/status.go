package mapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/mastodon/injection"
	"github.com/uakihir0/nostr-rest/server/mastodon/openapi"
)

func (h *MastodonHandler) GetApiV1AccountsUidStatuses(
	c echo.Context,
	uid string,
	params mopenapi.GetApiV1AccountsUidStatusesParams,
) error {
	statusService := minjection.StatusService()

	// TODO: for authentication fields
	_ = c.(*domain.Context).PubKey

	userPk := domain.UserPubKey(uid)
	options := params.ToTimeLineOptions()
	responses, err := statusService.GetUserStatues(userPk, options)
	if err != nil {
		return err
	}

	statuses := make([]mopenapi.Status, len(responses))
	for i, status := range responses {
		statuses[i] = ToStatus(status)
	}

	return c.JSON(
		http.StatusOK,
		statuses,
	)
}
