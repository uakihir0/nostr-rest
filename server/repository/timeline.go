package repository

import (
	"github.com/nbd-wtf/go-nostr"
	"github.com/samber/lo"

	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/util"
)

type RelayTimelineRepository struct {
}

var timelineRepositoryLock = util.Lock[RelayTimelineRepository]{}
var _ domain.TimelineRepository = (*RelayTimelineRepository)(nil)

// NewRelayTimelineRepository
// Create a new timeline repository
func NewRelayTimelineRepository() *RelayTimelineRepository {
	return timelineRepositoryLock.Once(
		func() *RelayTimelineRepository {
			return &RelayTimelineRepository{}
		},
	)
}

// GetTimelines
func (r *RelayTimelineRepository) GetTimelines(
	pks []domain.UserPubKey,
	options domain.PagingOptions,
) ([]domain.Timeline, error) {

	events := GetEventsByAuthor(
		[]int{
			nostr.KindTextNote,
			nostr.KindRepost,
		},
		pks,
		options,
	)

	// Distinct public keys
	pkMap := make(map[string]bool)
	posts := make([]domain.Timeline, 0)

	for _, event := range events {
		if !pkMap[event.Sig] {
			pkMap[event.Sig] = true

			if event.Kind == nostr.KindTextNote {
				post, err := MarshalPostEvent(event)
				if err != nil {
					return nil, err

				}

				posts = append(posts,
					domain.Timeline{
						Post: lo.ToPtr(post.ToPost()),
					},
				)
			}

			if event.Kind == nostr.KindRepost {
				repost, err := MarshalRepostEvent(event)
				if err != nil {
					return nil, err

				}

				posts = append(posts,
					domain.Timeline{
						Repost: lo.ToPtr(repost.ToRepost()),
					},
				)
			}
		}
	}

	return posts, nil
}
