package service

import (
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/util"
)

type UserService struct {
	userRepository domain.UserRepository
}

var userServiceLock = util.Lock[UserService]{}

func NewUserService(
	userRepository domain.UserRepository,
) *UserService {
	return userServiceLock.Once(
		func() *UserService {
			return &UserService{
				userRepository: userRepository,
			}
		},
	)
}

// GetUsers
// Get user information by public key
func (s *UserService) GetUsers(
	pks []domain.UserPubKey,
) ([]domain.User, error) {

	users, err := s.userRepository.GetUsers(pks)
	if err != nil {
		return nil, err
	}
	return users, nil
}
