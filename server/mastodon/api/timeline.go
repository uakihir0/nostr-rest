package mapi

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/uakihir0/nostr-rest/server/domain"
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

	return c.JSON(
		http.StatusOK,
		ToStatues(responses),
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
	timelineService := minjection.TimelineService()

	pk := c.(*domain.Context).PubKey
	options := params.ToTimeLineOptions()
	responses, err := timelineService.GetHomeTimeline(*pk, options)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		ToStatues(responses),
	)
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

