package service

import (
	"github.com/uakihir0/nostr-rest/server/domain"
	"sync"
)

var (
	once     sync.Once
	instance *UserService
)

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(
	userRepository domain.UserRepository,
) *UserService {
	once.Do(func() {
		instance = &UserService{
			userRepository: userRepository,
		}
	})
	return instance
}

// GetUsers
// Get user information by public key
func (s *UserService) GetUsers(
	pks []domain.UserPubKey,
) ([]*domain.User, error) {

	users, err := s.userRepository.GetUsers(pks)
	if err != nil {
		return nil, err
	}
	return users, nil
}
