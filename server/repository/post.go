package repository

import (
	"context"
	"github.com/nbd-wtf/go-nostr"
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/util"
	"time"
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

// SendPost
func (r *RelayPostRepository) SendPost(
	pk domain.UserPubKey,
	sk domain.UserSecretKey,
	text string,
) error {

	ev := nostr.Event{
		PubKey:    string(pk),
		CreatedAt: nostr.Timestamp(time.Now().Unix()),
		Kind:      1,
		Tags:      nil,
		Content:   text,
	}

	err := ev.Sign(string(sk))
	if err != nil {
		return err
	}

	SentEventAll(
		context.Background(),
		ev,
	)

	return nil
}

// GetPosts
func (r *RelayPostRepository) GetPosts(
	pks []domain.UserPubKey,
	maxResults int,
	sinceTime *time.Time,
	untilTime *time.Time,
) ([]domain.Post, error) {

	userPKs := lo.Map(pks,
		func(pk domain.UserPubKey, _ int) string {
			return string(pk)
		})

	filter := nostr.Filter{
		Kinds:   []int{1},
		Authors: userPKs,
		Limit:   maxResults,
	}

	if sinceTime != nil {
		filter.Since = lo.ToPtr(nostr.Timestamp(sinceTime.Unix()))
	}
	if untilTime != nil {
		filter.Until = lo.ToPtr(nostr.Timestamp(untilTime.Unix()))
	}

	events := QuerySyncAll(
		context.Background(),
		[]nostr.Filter{filter},
	)

	// Distinct public keys
	pkMap := make(map[string]bool)
	posts := make([]domain.Post, 0)

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
