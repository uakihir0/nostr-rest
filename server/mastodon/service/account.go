package mservice

import (
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/mastodon/domain"
	"github.com/uakihir0/nostr-rest/server/util"
	"time"
)

type AccountService struct {
	userRepository         domain.UserRepository
	postRepository         domain.PostRepository
	relationShipRepository domain.RelationShipRepository
}

var userServiceLock = util.Lock[AccountService]{}

func NewAccountService(
	userRepository domain.UserRepository,
	postRepository domain.PostRepository,
	relationShipRepository domain.RelationShipRepository,
) *AccountService {
	return userServiceLock.Once(
		func() *AccountService {
			return &AccountService{
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

	// Get user metadata first
	pks := []domain.UserPubKey{pk}
	users, err := s.userRepository.GetUsers(pks)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}

	user := users[0]
	acc := mdomain.Account{
		ID:          string(user.PubKey),
		Name:        user.Name,
		DisplayName: user.DisplayName,
		Picture:     user.Picture,
		Banner:      user.Banner,
		Website:     user.Website,
		About:       user.About,
		Lud06:       user.Lud06,
		CreatedAt:   time.Unix(user.CreatedAt, 0),
		LastStatsAt: nil,

		StatusesCount:  0,
		FollowingCount: 0,
		FollowersCount: 0,
	}

	statusCountCh := make(chan int)
	followingCountCh := make(chan int)
	followersCountCh := make(chan int)

	// Get user's posts count
	go func(ch chan int) {
		unix := time.Now().Unix()
		week := int64(60 * 60 * 24 * 7)
		posts, err := s.postRepository.GetPosts(
			[]domain.UserPubKey{user.PubKey}, 1000,
			lo.ToPtr(time.Unix(unix-week, 0)),
			lo.ToPtr(time.Unix(unix, 0)),
		)
		if err != nil {
			ch <- 0
			return
		}
		if len(posts) > 0 {
			// Set the last stats at here is a bit tricky
			acc.LastStatsAt = lo.ToPtr(posts[0].CreatedAt)
		}

		ch <- len(posts)
	}(statusCountCh)

	// Get user's following count
	go func(ch chan int) {
		following, err := s.relationShipRepository.GetFollowings(user.PubKey)
		if err != nil {
			ch <- 0
			return
		}
		ch <- len(following)
	}(followingCountCh)

	// Get user's followers count
	go func(ch chan int) {
		followers, err := s.relationShipRepository.GetFollowers(user.PubKey)
		if err != nil {
			ch <- 0
			return
		}
		ch <- len(followers)
	}(followersCountCh)

	acc.StatusesCount = <-statusCountCh
	acc.FollowingCount = <-followingCountCh
	acc.FollowersCount = <-followersCountCh
	return &acc, nil
}
