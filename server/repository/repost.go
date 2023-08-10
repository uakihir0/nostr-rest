package repository

import (
	"context"
	"github.com/nbd-wtf/go-nostr"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/util"
)

type RelayRepostRepository struct {
}

var repostRepositoryLock = util.Lock[RelayRepostRepository]{}
var _ domain.RepostRepository = (*RelayRepostRepository)(nil)

// NewRelayRepostRepository
// Create a new post repository
func NewRelayRepostRepository() *RelayRepostRepository {
	return repostRepositoryLock.Once(
		func() *RelayRepostRepository {
			return &RelayRepostRepository{}
		},
	)
}

func (r RelayRepostRepository) GetReposts(
	id domain.PostID,
) ([]domain.Repost, error) {

	events := QuerySyncAll(
		context.Background(),
		[]nostr.Filter{{
			Kinds: []int{nostr.KindRepost},
			Tags: map[string][]string{
				"e": {string(id)},
			},
			Limit: 1000,
		}},
	)

	pkMap := make(map[string]bool)
	results := make([]domain.Repost, 0)

	for _, event := range events {
		if !pkMap[event.Sig] {
			pkMap[event.Sig] = true

			re, err := MarshalRepostEvent(event)
			if err != nil {
				return nil, err
			}
			results = append(results, re.ToRepost())
		}
	}

	return results, nil
}
