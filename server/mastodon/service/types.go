package mservice

import (
	"github.com/uakihir0/nostr-rest/server/util"
	"time"

	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/mastodon/domain"
)

type TypeService struct {
	userRepository         domain.UserRepository
	postRepository         domain.PostRepository
	reactionRepository     domain.ReactionRepository
	relationShipRepository domain.RelationShipRepository
}

var typeServiceLock = util.Lock[TypeService]{}

func NewTypeService(
	userRepository domain.UserRepository,
	postRepository domain.PostRepository,
	reactionRepository domain.ReactionRepository,
	relationShipRepository domain.RelationShipRepository,
) *TypeService {
	return typeServiceLock.Once(
		func() *TypeService {
			return &TypeService{
				userRepository:         userRepository,
				postRepository:         postRepository,
				reactionRepository:     reactionRepository,
				relationShipRepository: relationShipRepository,
			}
		},
	)
}

// --------------------------------------------------------------------- //
// ACCOUNT
// --------------------------------------------------------------------- //

// AccountID
// make mastodon account domain object.
func (s *TypeService) AccountID(
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

	return s.Account(*users[0])
}

// Account
// make mastodon account domain object.
func (s *TypeService) Account(
	user domain.User,
) (*mdomain.Account, error) {

	acc := &mdomain.Account{
		ID:          mdomain.AccountID(user.PubKey),
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
	return acc, nil
}

// --------------------------------------------------------------------- //
// STATUS
// --------------------------------------------------------------------- //

// Status
// make mastodon status domain object.
func (s *TypeService) Status(
	post domain.Post,
) (*mdomain.Status, error) {

	acc, err := s.AccountID(post.UserPubKey)
	if err != nil {
		return nil, err
	}

	status := &mdomain.Status{
		ID: mdomain.NewStatusID(
			string(post.ID),
			post.CreatedAt,
		),
		Text:      post.Content,
		Account:   *acc,
		CreatedAt: post.CreatedAt,

		FavouritesCount: 0,
		ReblogsCount:    0,
	}

	favouritesCountCh := make(chan int)
	reblogsCountCh := make(chan int)

	// Get post's reactions count
	go func(ch chan int) {
		reactions, err := s.reactionRepository.GetReactions(post.ID)
		if err != nil {
			ch <- 0
			return
		}
		ch <- len(reactions)
	}(favouritesCountCh)

	// Get post's repost count
	go func(ch chan int) {
		reactions, err := s.reactionRepository.GetReactions(post.ID)
		if err != nil {
			ch <- 0
			return
		}
		ch <- len(reactions)
	}(reblogsCountCh)

	status.FavouritesCount = <-favouritesCountCh
	status.ReblogsCount = <-reblogsCountCh
	return status, nil
}

// Statuses
// make mastodon status domain object.
func (s *TypeService) Statuses(
	posts []domain.Post,
) ([]mdomain.Status, error) {

	statuses := make([]mdomain.Status, len(posts))
	for i, post := range posts {
		status, err := s.Status(post)
		if err != nil {
			return nil, err
		}
		statuses[i] = *status
	}
	return statuses, nil
}
