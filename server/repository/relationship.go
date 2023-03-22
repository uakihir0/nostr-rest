package repository

import (
	"context"
	"github.com/nbd-wtf/go-nostr"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/util"
)

type RelayRelationShipRepository struct {
}

var relationShipRepositoryLock = util.Lock[RelayRelationShipRepository]{}
var _ domain.RelationShipRepository = (*RelayRelationShipRepository)(nil)

// NewRelayRelationShipRepository
// Create a new relay relationship repository
func NewRelayRelationShipRepository() *RelayRelationShipRepository {
	return relationShipRepositoryLock.Once(
		func() *RelayRelationShipRepository {
			return &RelayRelationShipRepository{}
		},
	)
}

// GetFollowings
// Get public keys of users specific user is following
func (r *RelayRelationShipRepository) GetFollowings(
	pk domain.UserPubKey,
) ([]domain.UserPubKey, error) {

	events := QuerySyncAll(
		context.Background(),
		[]nostr.Filter{{
			Kinds:   []int{3},
			Authors: []string{string(pk)},
		}},
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

// GetFollowers
// Get public keys of users specific user is followed by
func (r *RelayRelationShipRepository) GetFollowers(
	pk domain.UserPubKey,
) ([]domain.UserPubKey, error) {

	options := DefaultQueryOptions()
	options.TimeoutSeconds = 3

	events := QuerySyncAllWithOptions(
		context.Background(),
		[]nostr.Filter{{
			Kinds: []int{3},
			Tags: map[string][]string{
				"p": {string(pk)},
			},
		}},
		options,
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
