package api

import (
	"net/http"

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

	var maxResults = 20
	if params.MaxResults != nil {
		maxResults = *params.MaxResults
	}

	// Get following user's post as timeline
	posts, err := postService.GetPosts(
		followingPks,
		maxResults,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	// Map by public key
	postsPks := lo.Map(posts,
		func(p *domain.Post, _ int) domain.UserPubKey {
			return p.UserPubKey
		})

	// Add specified user's public key
	postsPks = append(postsPks, myPk)

	// Distinct by user public key
	postsPks = lo.FindDuplicatesBy(postsPks,
		func(pk domain.UserPubKey) string {
			return string(pk)
		})

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
func (h *Handler) GetV1TimelinesUser(ctx echo.Context, params openapi.GetV1TimelinesUserParams) error {
	//TODO implement me
	panic("implement me")
}
