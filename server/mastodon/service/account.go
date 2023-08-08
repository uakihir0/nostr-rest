package mservice

import (
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/mastodon/domain"
	"github.com/uakihir0/nostr-rest/server/util"
)

type AccountService struct {
	typeService            *TypeService
	userRepository         domain.UserRepository
	postRepository         domain.PostRepository
	relationShipRepository domain.RelationShipRepository
}

var userServiceLock = util.Lock[AccountService]{}

func NewAccountService(
	typeService *TypeService,
	userRepository domain.UserRepository,
	postRepository domain.PostRepository,
	relationShipRepository domain.RelationShipRepository,
) *AccountService {
	return userServiceLock.Once(
		func() *AccountService {
			return &AccountService{
				typeService:            typeService,
				userRepository:         userRepository,
				postRepository:         postRepository,
				relationShipRepository: relationShipRepository,
			}
		},
	)
}

// GetAccount
func (s *AccountService) GetAccount(
	pk domain.UserPubKey,
) (*mdomain.Account, error) {
	return s.typeService.AccountID(pk)
}
