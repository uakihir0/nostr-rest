package mservice

import (
	"github.com/uakihir0/nostr-rest/server/domain"
	mdomain "github.com/uakihir0/nostr-rest/server/mastodon/domain"
	"github.com/uakihir0/nostr-rest/server/util"
)

type TimelineService struct {
	typeService            *TypeService
	userRepository         domain.UserRepository
	postRepository         domain.PostRepository
	relationShipRepository domain.RelationShipRepository
}

var timelineServiceLock = util.Lock[TimelineService]{}

func NewTimelineService(
	typeService *TypeService,
	userRepository domain.UserRepository,
	postRepository domain.PostRepository,
	relationShipRepository domain.RelationShipRepository,
) *TimelineService {
	return timelineServiceLock.Once(
		func() *TimelineService {
			return &TimelineService{
				typeService:            typeService,
				userRepository:         userRepository,
				postRepository:         postRepository,
				relationShipRepository: relationShipRepository,
			}
		},
	)
}

// GetPublicTimeline
func (s *TimelineService) GetPublicTimeline(
	op mdomain.TimelineOptions,
) ([]mdomain.Status, error) {

	posts, err := s.postRepository.GetPublicPosts(
		op.GetLimit(20),
		op.GetSinceTime(),
		op.GetUntilTime(),
	)
	if err != nil {
		return nil, err
	}

	return s.typeService.Statuses(posts)
}
