package repository

import (
	"context"
	"fmt"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/nbd-wtf/go-nostr"
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
)

type UserEventCache struct {
	Event     *UserEvent
	ExpiredAt int64
}

type RelayUserRepository struct {
	Cache *lru.Cache[domain.UserPubKey, *UserEventCache]
}

var _ domain.UserRepository = (*RelayUserRepository)(nil)

// NewRelayUserRepository
// Create a new relay user repository
func NewRelayUserRepository() *RelayUserRepository {
	cache, err := lru.New[domain.UserPubKey, *UserEventCache](200)
	if err != nil {
		panic("Error on NewRelayUserRepository Init")
	}

	return &RelayUserRepository{
		Cache: cache,
	}
}

// GetUserFromCache
// Retrieve user information from cached data
func (r *RelayUserRepository) GetUserFromCache(
	pk domain.UserPubKey,
) *domain.User {
	value, ok := r.Cache.Get(pk)
	if !ok {
		return nil
	}

	now := time.Now().Unix()
	if value.ExpiredAt <= now {
		r.Cache.Remove(pk)
		return nil
	}

	return value.Event.ToUser()
}

func (r *RelayUserRepository) SetUserCache(
	pk domain.UserPubKey,
	event *UserEvent,
) {
	r.Cache.Add(pk,
		&UserEventCache{
			Event:     event,
			ExpiredAt: time.Now().Unix() + (60 * 15),
		})
}

// GetUsers
// Retrieve user information
func (r *RelayUserRepository) GetUsers(
	pks []domain.UserPubKey,
) ([]*domain.User, error) {

	// Temporary storage of acquired user information
	userMap := make(map[domain.UserPubKey]*domain.User)
	//　UserPKs that does not exist in cache
	nonUserPKs := make([]domain.UserPubKey, 0)

	for _, pk := range pks {

		// Fetch data from cache first
		user := r.GetUserFromCache(pk)
		if user != nil {
			// Use data in cache
			userMap[pk] = user
		} else {
			// Record if not available from cache
			nonUserPKs = append(nonUserPKs, pk)
		}
	}

	if len(nonUserPKs) > 0 {
		userPKs := lo.Map(nonUserPKs,
			func(pk domain.UserPubKey, _ int) string {
				return string(pk)
			})

		events := QuerySyncAllWithGuard(
			context.Background(),
			[]nostr.Filter{{
				Kinds:   []int{0},
				Authors: userPKs,
			}},
			len(userPKs),
		)

		for _, event := range events {
			ue, err := MarshalUserEvent(event)
			if err != nil {
				return nil, err
			}

			pk := domain.UserPubKey(event.PubKey)
			userMap[pk] = ue.ToUser()
			r.SetUserCache(pk, ue)
		}
	}

	// Sort by request order
	result := make([]*domain.User, 0)
	for _, pk := range pks {
		user, ok := userMap[pk]
		if !ok {
			return nil, fmt.Errorf("not found user data: PK=%s", pk)
		}
		// Append user if present
		result = append(result, user)
	}

	return result, nil
}
