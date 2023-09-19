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
	repostRepository       domain.RepostRepository
	timelineRepository     domain.TimelineRepository
	relationShipRepository domain.RelationShipRepository
}

var timelineServiceLock = util.Lock[TimelineService]{}

func NewTimelineService(
	typeService *TypeService,
	userRepository domain.UserRepository,
	postRepository domain.PostRepository,
	repostRepository domain.RepostRepository,
	timelineRepository domain.TimelineRepository,
	relationShipRepository domain.RelationShipRepository,
) *TimelineService {
	return timelineServiceLock.Once(
		func() *TimelineService {
			return &TimelineService{
				typeService:            typeService,
				userRepository:         userRepository,
				postRepository:         postRepository,
				repostRepository:       repostRepository,
				timelineRepository:     timelineRepository,
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
		op.ToPagingOptions(20),
	)
	if err != nil {
		return nil, err
	}

	return s.typeService.Statuses(posts)
}

// GetHomeTimeline
func (s *TimelineService) GetHomeTimeline(
	pk domain.UserPubKey,
	op mdomain.TimelineOptions,
) ([]mdomain.Status, error) {

	// フォロワーリストを取得
	followers, err := s.relationShipRepository.GetFollowers(pk)
	if err != nil {
		return nil, err
	}

	// 自分自身の投稿も取得対象に含める
	followers = append(followers, pk)

	// フォロワーの投稿を取得
	posts, err := s.timelineRepository.GetTimelines(
		followers, op.ToPagingOptions(20),
	)
	if err != nil {
		return nil, err
	}

	return s.typeService.Timelines(posts)
}
