package repository

import (
	"context"
	"github.com/nbd-wtf/go-nostr"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/util"
)

type RelayRelationShipRepository struct {
	FollowingsCache StringCacheMap
	FollowersCache  StringCacheMap
}

var relationShipRepositoryLock = util.Lock[RelayRelationShipRepository]{}
var _ domain.RelationShipRepository = (*RelayRelationShipRepository)(nil)

// NewRelayRelationShipRepository
// Create a new relay relationship repository
func NewRelayRelationShipRepository() *RelayRelationShipRepository {
	return relationShipRepositoryLock.Once(
		func() *RelayRelationShipRepository {
			return &RelayRelationShipRepository{
				FollowingsCache: NewStringCacheMap(200),
				FollowersCache:  NewStringCacheMap(200),
			}
		},
	)
}

var followingsLimitMap = util.NewLimitMap(1)

// GetFollowings
// Get public keys of users specific user is following
func (r *RelayRelationShipRepository) GetFollowings(
	pk domain.UserPubKey,
) ([]domain.UserPubKey, error) {

	ctx := context.Background()
	followingsLimitMap.Acquire(ctx, string(pk))
	defer followingsLimitMap.Release(string(pk))

	// Get events from cache or query
	events := ManageEventsFromString(
		r.FollowingsCache, string(pk),
		func() []*nostr.Event {
			return QuerySyncAll(
				ctx,
				[]nostr.Filter{{
					Kinds:   []int{3},
					Authors: []string{string(pk)},
				}},
			)
		},
	)

	// Distinct user public keys
	pkMap := make(map[string]bool)
	pks := make([]domain.UserPubKey, 0)

	for _, event := range events {
		for _, tag := range event.Tags {
			if !pkMap[tag[1]] {
				pkMap[tag[1]] = true
				pks = append(pks, domain.UserPubKey(tag[1]))
			}
		}
	}

	return pks, nil
}

var followersLimitMap = util.NewLimitMap(1)

// GetFollowers
// Get public keys of users specific user is followed by
func (r *RelayRelationShipRepository) GetFollowers(
	pk domain.UserPubKey,
) ([]domain.UserPubKey, error) {

	ctx := context.Background()
	followersLimitMap.Acquire(ctx, string(pk))
	defer followersLimitMap.Release(string(pk))

	// Get events from cache or query
	events := ManageEventsFromString(
		r.FollowingsCache, string(pk),
		func() []*nostr.Event {
			return QuerySyncAll(
				ctx,
				[]nostr.Filter{{
					Kinds: []int{3},
					Tags: map[string][]string{
						"p": {string(pk)},
					},
				}},
			)
		},
	)

	// Distinct user public keys
	pkMap := make(map[string]bool)
	pks := make([]domain.UserPubKey, 0)

	for _, event := range events {
		if !pkMap[event.PubKey] {
			pkMap[event.PubKey] = true
			pks = append(pks, domain.UserPubKey(event.PubKey))
		}
	}

	return pks, nil
}
