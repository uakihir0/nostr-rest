package mservice

import (
	"github.com/uakihir0/nostr-rest/server/util"
	"sync"
	"time"

	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/mastodon/domain"
)

type TypeService struct {
	userRepository         domain.UserRepository
	postRepository         domain.PostRepository
	repostRepository       domain.RepostRepository
	reactionRepository     domain.ReactionRepository
	relationShipRepository domain.RelationShipRepository
}

var typeServiceLock = util.Lock[TypeService]{}

func NewTypeService(
	userRepository domain.UserRepository,
	postRepository domain.PostRepository,
	repostRepository domain.RepostRepository,
	reactionRepository domain.ReactionRepository,
	relationShipRepository domain.RelationShipRepository,
) *TypeService {
	return typeServiceLock.Once(
		func() *TypeService {
			return &TypeService{
				userRepository:         userRepository,
				postRepository:         postRepository,
				repostRepository:       repostRepository,
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

	return s.Account(users[0])
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

	var wg sync.WaitGroup
	wg.Add(3)

	// Get user's posts count
	go func(wg *sync.WaitGroup) {
		posts, err := s.postRepository.GetUserLatestPosts(user.PubKey)
		if err != nil {
			return
		}

		if len(posts) > 0 {
			// Set the last stats at here is a bit tricky
			acc.LastStatsAt = lo.ToPtr(posts[0].CreatedAt)
		}
		acc.StatusesCount = len(posts)
		wg.Done()
	}(&wg)

	// Get user's following count
	go func(wg *sync.WaitGroup) {
		following, err := s.relationShipRepository.GetFollowings(user.PubKey)
		if err != nil {
			return
		}
		acc.FollowingCount = len(following)
		wg.Done()
	}(&wg)

	// Get user's followers count
	go func(wg *sync.WaitGroup) {
		followers, err := s.relationShipRepository.GetFollowers(user.PubKey)
		if err != nil {
			return
		}
		acc.FollowersCount = len(followers)
		wg.Done()
	}(&wg)

	wg.Wait()
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

	status := &mdomain.Status{
		ID: mdomain.NewStatusID(
			string(post.ID),
			post.CreatedAt,
		),
		Text:      post.Content,
		CreatedAt: post.CreatedAt,

		FavouritesCount: 0,
		ReblogsCount:    0,
	}

	var wg sync.WaitGroup
	wg.Add(3)

	// Get post's account
	go func(wg *sync.WaitGroup) {
		acc, err := s.AccountID(post.UserPubKey)
		if err != nil {
			return
		}
		status.Account = *acc
		wg.Done()
	}(&wg)

	// Get post's reactions count
	go func(wg *sync.WaitGroup) {
		reactions, err := s.reactionRepository.GetReactions(post.ID)
		if err != nil {
			return
		}
		status.FavouritesCount = len(reactions)
		wg.Done()
	}(&wg)

	// Get post's repost count
	go func(wg *sync.WaitGroup) {
		reposts, err := s.repostRepository.GetReposts(post.ID)
		if err != nil {
			return
		}
		status.ReblogsCount = len(reposts)
		wg.Done()
	}(&wg)

	wg.Wait()
	return status, nil
}

// Statuses
// make mastodon status domain object.
func (s *TypeService) Statuses(
	posts []domain.Post,
) ([]mdomain.Status, error) {

	statuses := make([]mdomain.Status, len(posts))
	errors := make([]error, len(posts))

	var wg sync.WaitGroup
	wg.Add(len(posts))

	// Get statuses concurrently
	for i, post := range posts {
		go func(i int, post domain.Post, wg *sync.WaitGroup) {
			status, err := s.Status(post)
			if err != nil {
				errors[i] = err
				wg.Done()
				return
			}
			statuses[i] = *status
			wg.Done()
		}(i, post, &wg)
	}

	wg.Wait()

	// Check if there is any error
	for _, err := range errors {
		if err != nil {
			return nil, err
		}
	}

	return statuses, nil
}
