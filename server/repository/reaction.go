package repository

import (
	"context"
	"github.com/nbd-wtf/go-nostr"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/util"
)

type RelayReactionRepository struct {
}

var reactionRepositoryLock = util.Lock[RelayReactionRepository]{}
var _ domain.ReactionRepository = (*RelayReactionRepository)(nil)

// NewRelayReactionRepository
// Create a new post repository
func NewRelayReactionRepository() *RelayReactionRepository {
	return reactionRepositoryLock.Once(
		func() *RelayReactionRepository {
			return &RelayReactionRepository{}
		},
	)
}

func (r RelayReactionRepository) GetReactions(
	id domain.PostID,
) ([]domain.Reaction, error) {

	events := QuerySyncAll(
		context.Background(),
		[]nostr.Filter{{
			Kinds: []int{nostr.KindReaction},
			Tags: map[string][]string{
				"#e": {string(id)},
			},
			Limit: 1000,
		}},
	)

	results := make([]domain.Reaction, 0)
	for _, event := range events {
		re, err := MarshalReactionEvent(event)
		if err != nil {
			return nil, err
		}

		result := re.ToReaction()
		results = append(results, result)
	}

	return results, nil
}
