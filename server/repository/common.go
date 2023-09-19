package repository

import (
	"context"

	"github.com/nbd-wtf/go-nostr"
	"github.com/samber/lo"

	"github.com/uakihir0/nostr-rest/server/domain"
)

func GetEventsByAuthor(
	kinds []int,
	authors []domain.UserPubKey,
	options domain.PagingOptions,
) []*nostr.Event {

	authorPKs := lo.Map(authors,
		func(pk domain.UserPubKey, _ int) string {
			return string(pk)
		})

	filter := nostr.Filter{
		Kinds: kinds,
		Limit: options.MaxResults,
	}

	if len(authorPKs) > 0 {
		filter.Authors = authorPKs
	}
	if options.SinceTime != nil {
		filter.Since = lo.ToPtr(nostr.Timestamp(
			options.SinceTime.Unix()))
	}
	if options.UntilTime != nil {
		filter.Until = lo.ToPtr(nostr.Timestamp(
			options.UntilTime.Unix()))
	}

	return QuerySyncAll(
		context.Background(),
		[]nostr.Filter{filter},
	)
}
