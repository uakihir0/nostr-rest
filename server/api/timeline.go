package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/injection"
	"github.com/uakihir0/nostr-rest/server/openapi"
)

// GetV1TimelinesHome
func (h *Handler) GetV1TimelinesHome(c echo.Context, params openapi.GetV1TimelinesHomeParams) error {
	relationShipService := injection.RelationShipService()
	postService := injection.PostService()
	userService := injection.UserService()

	myPk := domain.UserPubKey(params.Pubkey)

	// Get user following user's public keys
	followingPks, err := relationShipService.GetFollowingPubKeys(myPk)
	if err != nil {
		return err
	}

	maxResults, sinceTime, untilTime, err := paging(
		params.MaxResults,
		params.UntilTime,
		params.SinceTime,
	)
	if err != nil {
		return err
	}

	// Get following user's post as timeline
	posts, err := postService.GetPosts(
		followingPks,
		maxResults,
		sinceTime,
		untilTime,
	)
	if err != nil {
		return err
	}

	postsPks := timelineUsers(posts,
		[]domain.UserPubKey{myPk},
	)

	users, err := userService.GetUsers(postsPks)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		ToTimeline(
			postsPks,
			posts,
			users,
		),
	)
}

// GetV1TimelinesUser
func (h *Handler) GetV1TimelinesUser(c echo.Context, params openapi.GetV1TimelinesUserParams) error {
	postService := injection.PostService()
	userService := injection.UserService()

	myPk := domain.UserPubKey(params.Pubkey)

	maxResults, sinceTime, untilTime, err := paging(
		params.MaxResults,
		params.UntilTime,
		params.SinceTime,
	)
	if err != nil {
		return err
	}

	// Get my post as timeline
	posts, err := postService.GetPosts(
		[]domain.UserPubKey{myPk},
		maxResults,
		sinceTime,
		untilTime,
	)
	if err != nil {
		return err
	}

	postsPks := timelineUsers(posts,
		[]domain.UserPubKey{myPk},
	)

	users, err := userService.GetUsers(postsPks)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		ToTimeline(
			postsPks,
			posts,
			users,
		),
	)
}

func paging(
	pMaxResults *openapi.MaxResultsParameter,
	pSinceTime *openapi.SinceTimeParameter,
	pUntilTime *openapi.UntilTimePatameter,
) (int, *time.Time, *time.Time, error) {

	var maxResults = 20
	if pMaxResults != nil {
		maxResults = *pMaxResults
	}

	var sinceTime *time.Time = nil
	if pSinceTime != nil {
		parse, err := time.Parse(TimeLayout, *pSinceTime)
		if err != nil {
			return 0, nil, nil, err
		}
		sinceTime = &parse
	}

	var untilTime *time.Time = nil
	if pUntilTime != nil {
		parse, err := time.Parse(TimeLayout, *pUntilTime)
		if err != nil {
			return 0, nil, nil, err
		}
		untilTime = &parse
	}

	return maxResults, sinceTime, untilTime, nil
}

func timelineUsers(
	posts []*domain.Post,
	additionalPks []domain.UserPubKey,
) []domain.UserPubKey {

	// Map by public key
	postsPks := lo.Map(posts,
		func(p *domain.Post, _ int) domain.UserPubKey {
			return p.UserPubKey
		})

	// Add additional public key
	for _, pk := range additionalPks {
		postsPks = append(postsPks, pk)
	}

	// Distinct by user public key
	postsPks = lo.FindDuplicatesBy(postsPks,
		func(pk domain.UserPubKey) string {
			return string(pk)
		})

	return postsPks
}
