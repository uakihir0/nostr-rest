package mservice

import (
	"github.com/uakihir0/nostr-rest/server/domain"
	mdomain "github.com/uakihir0/nostr-rest/server/mastodon/domain"
	"github.com/uakihir0/nostr-rest/server/util"
)

type StatusService struct {
	typeService            *TypeService
	userRepository         domain.UserRepository
	postRepository         domain.PostRepository
	relationShipRepository domain.RelationShipRepository
}

var statusServiceLock = util.Lock[StatusService]{}

func NewStatusService(
	typeService *TypeService,
	userRepository domain.UserRepository,
	postRepository domain.PostRepository,
	relationShipRepository domain.RelationShipRepository,
) *StatusService {
	return statusServiceLock.Once(
		func() *StatusService {
			return &StatusService{
				typeService:            typeService,
				userRepository:         userRepository,
				postRepository:         postRepository,
				relationShipRepository: relationShipRepository,
			}
		},
	)
}

// GetUserStatues
func (s *StatusService) GetUserStatues(
	pk domain.UserPubKey,
	op mdomain.TimelineOptions,
) ([]mdomain.Status, error) {

	limit := op.GetLimit(20)
	sinceTime := op.GetSinceTime()
	untilTime := op.GetUntilTime()

	go func() {
		// Get user info for store cache (for speed)
		_, _ = s.userRepository.GetUsers([]domain.UserPubKey{pk})
	}()

	// Get user metadata first
	pks := []domain.UserPubKey{pk}
	posts, err := s.postRepository.GetPosts(pks, limit, sinceTime, untilTime)
	if err != nil {
		return nil, err
	}

	return s.typeService.Statuses(posts)
}
