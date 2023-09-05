package mapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/uakihir0/nostr-rest/server/mastodon/injection"
	"github.com/uakihir0/nostr-rest/server/mastodon/openapi"
)

func (h *MastodonHandler) GetApiV1TimelinesPublic(
	c echo.Context,
	params mopenapi.GetApiV1TimelinesPublicParams,
) error {
	timelineService := minjection.TimelineService()

	options := params.ToTimeLineOptions()
	responses, err := timelineService.GetPublicTimeline(options)
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
