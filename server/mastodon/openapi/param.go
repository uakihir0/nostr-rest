package mopenapi

import (
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/mastodon/domain"
)

func (p GetApiV1AccountsUidStatusesParams) ToTimeLineOptions() mdomain.TimelineOptions {
	options := mdomain.TimelineOptions{
		Limit:          p.Limit,
		OnlyMedia:      p.OnlyMedia,
		ExcludeReblogs: p.ExcludeReblogs,
		ExcludeReplies: p.ExcludeReplies,
		Pinned:         p.Pinned,
		Tagged:         p.Tagged,
	}

	if p.MaxId != nil {
		options.MaxId = lo.ToPtr(mdomain.StatusID(*p.MaxId))
	}
	if p.SinceId != nil {
		options.SinceId = lo.ToPtr(mdomain.StatusID(*p.SinceId))
	}
	if p.MinId != nil {
		options.MinId = lo.ToPtr(mdomain.StatusID(*p.MinId))
	}
	return options
}

func (p GetApiV1TimelinesPublicParams) ToTimeLineOptions() mdomain.TimelineOptions {
	options := mdomain.TimelineOptions{
		Limit:          p.Limit,
		OnlyMedia:      p.OnlyMedia,
		ExcludeReblogs: lo.ToPtr(false),
		ExcludeReplies: lo.ToPtr(false),
		Pinned:         lo.ToPtr(false),
		Tagged:         nil,
	}

	if p.MaxId != nil {
		options.MaxId = lo.ToPtr(mdomain.StatusID(*p.MaxId))
	}
	if p.SinceId != nil {
		options.SinceId = lo.ToPtr(mdomain.StatusID(*p.SinceId))
	}
	if p.MinId != nil {
		options.MinId = lo.ToPtr(mdomain.StatusID(*p.MinId))
	}
	return options
}


func (p GetApiV1TimelinesHomeParams) ToTimeLineOptions() mdomain.TimelineOptions {
	options := mdomain.TimelineOptions{
		Limit:          p.Limit,
		OnlyMedia:      lo.ToPtr(false),
		ExcludeReblogs: lo.ToPtr(false),
		ExcludeReplies: lo.ToPtr(false),
		Pinned:         lo.ToPtr(false),
		Tagged:         nil,
	}

	if p.MaxId != nil {
		options.MaxId = lo.ToPtr(mdomain.StatusID(*p.MaxId))
	}
	if p.SinceId != nil {
		options.SinceId = lo.ToPtr(mdomain.StatusID(*p.SinceId))
	}
	if p.MinId != nil {
		options.MinId = lo.ToPtr(mdomain.StatusID(*p.MinId))
	}
	return options
}

