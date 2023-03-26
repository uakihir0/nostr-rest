package service

import (
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/util"
	"sort"
	"time"
)

type PostService struct {
	postRepository domain.PostRepository
}

var postServiceLock = util.Lock[PostService]{}

func NewPostService(
	postRepository domain.PostRepository,
) *PostService {
	return postServiceLock.Once(
		func() *PostService {
			return &PostService{
				postRepository: postRepository,
			}
		},
	)
}

// SendPost
func (s *PostService) SendPost(
	pk domain.UserPubKey,
	sk domain.UserSecretKey,
	text string,
) error {

	err := s.postRepository.SendPost(
		pk, sk, text,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetPosts
// Results return in descending order of creation time.
func (s *PostService) GetPosts(
	pks []domain.UserPubKey,
	maxResults int,
	sinceTime *time.Time,
	untilTime *time.Time,
) ([]*domain.Post, error) {

	posts, err := s.postRepository.GetPosts(
		pks, maxResults, sinceTime, untilTime,
	)
	if err != nil {
		return nil, err
	}

	sort.Slice(posts, func(i, j int) bool {
		// Sort by creation time in descending order of time
		return posts[i].CreatedAt.Unix() > posts[j].CreatedAt.Unix()
	})

	return lo.Subset(
		posts,
		0,
		uint(maxResults),
	), nil
}
