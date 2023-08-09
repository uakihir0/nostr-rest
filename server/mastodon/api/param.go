package mapi

import (
	"github.com/samber/lo"
	mdomain "github.com/uakihir0/nostr-rest/server/mastodon/domain"
	mopenapi "github.com/uakihir0/nostr-rest/server/mastodon/openapi"
)

func ToTimeLineOptions(
	params mopenapi.GetApiV1AccountsUidStatusesParams,
) mdomain.TimelineOptions {

	options := mdomain.TimelineOptions{}

	if params.MaxId != nil {
		options.MaxId = lo.ToPtr(mdomain.StatusID(*params.MaxId))
	}
	if params.MaxId != nil {
		options.SinceId = lo.ToPtr(mdomain.StatusID(*params.SinceId))
	}
	if params.MaxId != nil {
		options.MinId = lo.ToPtr(mdomain.StatusID(*params.MinId))
	}

	options.Limit = params.Limit
	options.OnlyMedia = params.OnlyMedia
	options.ExcludeReblogs = params.ExcludeReblogs
	options.ExcludeReplies = params.ExcludeReplies
	options.Pinned = params.Pinned
	options.Tagged = params.Tagged

	return options
}
