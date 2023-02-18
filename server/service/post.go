package service

import (
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/util"
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

// GetPosts
func (s *PostService) GetPosts(
	pks []domain.UserPubKey,
	maxResults int,
	startTime *time.Time,
	endTime *time.Time,
) ([]*domain.Post, error) {

	posts, err := s.postRepository.GetPosts(
		pks,
		maxResults,
		startTime,
		endTime,
	)

	if err != nil {
		return nil, err
	}
	return posts, nil
}
