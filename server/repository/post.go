package repository

import (
	"context"
	"github.com/nbd-wtf/go-nostr"
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/util"
)

type RelayPostRepository struct {
}

var postRepositoryLock = util.Lock[RelayPostRepository]{}
var _ domain.PostRepository = (*RelayPostRepository)(nil)

// NewRelayPostRepository
// Create a new post repository
func NewRelayPostRepository() *RelayPostRepository {
	return postRepositoryLock.Once(
		func() *RelayPostRepository {
			return &RelayPostRepository{}
		},
	)
}

// GetPosts
func (r *RelayPostRepository) GetPosts(pks []domain.UserPubKey) ([]*domain.Post, error) {

	userPKs := lo.Map(pks,
		func(pk domain.UserPubKey, _ int) string {
			return string(pk)
		})

	events := QuerySyncAll(
		context.Background(),
		[]nostr.Filter{{
			Kinds:   []int{1},
			Authors: userPKs,
		}},
	)

	// Distinct public keys
	pkMap := make(map[string]bool)
	posts := make([]*domain.Post, 0)

	for _, event := range events {
		if !pkMap[event.Sig] {
			pkMap[event.Sig] = true
			post, err := MarshalPostEvent(event)
			if err != nil {
				return nil, err

			}

			posts = append(posts, post.ToPost())
		}
	}

	return posts, nil
}
