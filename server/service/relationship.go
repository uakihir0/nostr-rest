package service

import (
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/util"
)

type RelationShipService struct {
	relationShipRepository domain.RelationShipRepository
}

var relationShipServiceLock = util.Lock[RelationShipService]{}

func NewRelationShipService(
	relationShipRepository domain.RelationShipRepository,
) *RelationShipService {
	return relationShipServiceLock.Once(
		func() *RelationShipService {
			return &RelationShipService{
				relationShipRepository: relationShipRepository,
			}
		},
	)
}

// GetFollowingPubKeys
// Get the public key of a user who is followed by a specific user.
func (s *RelationShipService) GetFollowingPubKeys(pk domain.UserPubKey) ([]domain.UserPubKey, error) {

	followings, err := s.relationShipRepository.GetFollowings(pk)
	if err != nil {
		return nil, err
	}
	return followings, nil
}

// GetFollowersPubKeys
// Get the public key of a user who follows a specific user.
func (s *RelationShipService) GetFollowersPubKeys(pk domain.UserPubKey) ([]domain.UserPubKey, error) {

	followings, err := s.relationShipRepository.GetFollowers(pk)
	if err != nil {
		return nil, err
	}
	return followings, nil
}
