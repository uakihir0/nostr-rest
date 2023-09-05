package mapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/uakihir0/nostr-rest/server/mastodon/injection"
	"github.com/uakihir0/nostr-rest/server/mastodon/openapi"
)

// GetApiV1TimelinesPublic
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

// GetApiV1TimelinesTagHashtag
func (h *MastodonHandler) GetApiV1TimelinesTagHashtag(
	c echo.Context,
	hashtag mopenapi.HashTagPathParam,
	params mopenapi.GetApiV1TimelinesTagHashtagParams,
) error {
	//TODO implement me
	panic("implement me")
}

// GetApiV1TimelinesHome
func (h *MastodonHandler) GetApiV1TimelinesHome(
	c echo.Context,
	params mopenapi.GetApiV1TimelinesHomeParams,
) error {
	//TODO implement me
	panic("implement me")
}

// GetApiV1TimelinesListListId
func (h *MastodonHandler) GetApiV1TimelinesListListId(
	ctx echo.Context,
	lid string,
	params mopenapi.GetApiV1TimelinesListListIdParams,
) error {
	//TODO implement me
	panic("implement me")
}
