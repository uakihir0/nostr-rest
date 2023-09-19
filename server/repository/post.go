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
	LatestPostCache StringCacheMap
	GetRepliesCache StringCacheMap
}

var postRepositoryLock = util.Lock[RelayPostRepository]{}
var _ domain.PostRepository = (*RelayPostRepository)(nil)

// NewRelayPostRepository
// Create a new post repository
func NewRelayPostRepository() *RelayPostRepository {
	return postRepositoryLock.Once(
		func() *RelayPostRepository {
			return &RelayPostRepository{
				LatestPostCache: NewStringCacheMap(200),
				GetRepliesCache: NewStringCacheMap(200),
			}
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

var getPostLimitMap = util.NewLimitMap(1)

// GetPost
func (r *RelayPostRepository) GetPost(
	id domain.PostID,
) (*domain.Post, error) {

	ctx := context.Background()
	getPostLimitMap.Acquire(ctx, string(id))
	defer getPostLimitMap.Release(string(id))

	// Get events from cache or query
	events := ManageEventsFromString(
		r.GetRepliesCache, string(id),
		func() []*nostr.Event {
			filter := nostr.Filter{
				Kinds: []int{1},
				IDs:   []string{string(id)},
				Limit: 1,
			}
			return QuerySyncAll(
				context.Background(),
				[]nostr.Filter{filter},
			)
		},
	)

	if len(events) > 0 {
		post, err := MarshalPostEvent(events[0])
		if err != nil {
			return nil, err

		}
		return lo.ToPtr(post.ToPost()), nil
	}

	return nil, nil
}

// GetPosts
func (r *RelayPostRepository) GetPosts(
	pks []domain.UserPubKey,
	options domain.PagingOptions,
) ([]domain.Post, error) {

	events := GetEventsByAuthor(
		[]int{nostr.KindTextNote},
		pks, options,
	)
	return eventsToPosts(events)
}

// GetPublicPosts
func (r *RelayPostRepository) GetPublicPosts(
	options domain.PagingOptions,
) ([]domain.Post, error) {

	return r.GetPosts(
		[]domain.UserPubKey{},
		options,
	)
}

var latestPostsLimitMap = util.NewLimitMap(1)

// GetUserLatestPosts
func (r *RelayPostRepository) GetUserLatestPosts(
	pk domain.UserPubKey,
) ([]domain.Post, error) {

	ctx := context.Background()
	latestPostsLimitMap.Acquire(ctx, string(pk))
	defer latestPostsLimitMap.Release(string(pk))

	// Get events from cache or query
	events := ManageEventsFromString(
		r.LatestPostCache, string(pk),
		func() []*nostr.Event {
			filter := nostr.Filter{
				Kinds:   []int{1},
				Authors: []string{string(pk)},
				Limit:   1000,
			}

			unix := time.Now().Unix()
			week := int64(60 * 60 * 24 * 7)
			filter.Since = lo.ToPtr(nostr.Timestamp(unix - week))
			filter.Until = lo.ToPtr(nostr.Timestamp(unix))

			return QuerySyncAll(
				context.Background(),
				[]nostr.Filter{filter},
			)
		},
	)

	return eventsToPosts(events)
}

var getRepliesLimitMap = util.NewLimitMap(1)

// GetReplies
func (r *RelayPostRepository) GetReplies(
	id domain.PostID,
) ([]domain.Post, error) {

	ctx := context.Background()
	latestPostsLimitMap.Acquire(ctx, string(id))
	defer latestPostsLimitMap.Release(string(id))

	// Get events from cache or query
	events := ManageEventsFromString(
		r.LatestPostCache, string(id),
		func() []*nostr.Event {
			filter := nostr.Filter{
				Kinds: []int{1},
				Tags: map[string][]string{
					"e": {string(id)},
				},
				Limit: 1000,
			}
			return QuerySyncAll(
				context.Background(),
				[]nostr.Filter{filter},
			)
		},
	)

	return eventsToPosts(events)
}

func eventsToPosts(
	events []*nostr.Event,
) ([]domain.Post, error) {

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
