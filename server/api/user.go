package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/injection"
	"github.com/uakihir0/nostr-rest/server/openapi"
)

// GetV1Users
// GET Users Profiles
func (h *SimpleHandler) GetV1Users(c echo.Context, params openapi.GetV1UsersParams) error {
	userService := injection.UserService()

	pks := []domain.UserPubKey{
		domain.UserPubKey(params.Pubkey),
	}

	users, err := userService.GetUsers(pks)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		ToUser(users[0]),
	)
}

// PostV1Users
// GET User Profiles
func (h *SimpleHandler) PostV1Users(c echo.Context) error {
	userService := injection.UserService()

	request := new(openapi.UsersPubKeyRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	pks := lo.Map(request.Pubkeys,
		func(pk string, _ int) domain.UserPubKey {
			return domain.UserPubKey(pk)
		})

	users, err := userService.GetUsers(pks)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		ToUsersResponse(users),
	)
}

// GetV1UsersFollowing
// Get following users
func (h *SimpleHandler) GetV1UsersFollowing(c echo.Context, params openapi.GetV1UsersFollowingParams) error {
	userService := injection.UserService()
	relationShipService := injection.RelationShipService()

	pk := domain.UserPubKey(params.Pubkey)

	// Get public keys first
	pks, err := relationShipService.GetFollowingPubKeys(pk)
	if err != nil {
		return err
	}

	// Get user objects from public keys
	users, err := userService.GetUsers(pks)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		ToUsersResponse(users),
	)
}

// GetV1UsersFollowingPubkeys
// Get User's following public keys
func (h *SimpleHandler) GetV1UsersFollowingPubkeys(c echo.Context, params openapi.GetV1UsersFollowingPubkeysParams) error {
	relationShipService := injection.RelationShipService()

	pk := domain.UserPubKey(params.Pubkey)

	pks, err := relationShipService.GetFollowingPubKeys(pk)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		ToPubKeysResponse(pks),
	)
}

// GetV1UsersFollowers
// Get user's followers
func (h *SimpleHandler) GetV1UsersFollowers(c echo.Context, params openapi.GetV1UsersFollowersParams) error {
	userService := injection.UserService()
	relationShipService := injection.RelationShipService()

	pk := domain.UserPubKey(params.Pubkey)

	// Get public keys first
	pks, err := relationShipService.GetFollowersPubKeys(pk)
	if err != nil {
		return err
	}

	// Get user objects from public keys
	users, err := userService.GetUsers(pks)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		ToUsersResponse(users),
	)
}

// GetV1UsersFollowersPubkeys
// Get user's follower's public keys
func (h *SimpleHandler) GetV1UsersFollowersPubkeys(c echo.Context, params openapi.GetV1UsersFollowersPubkeysParams) error {
	relationShipService := injection.RelationShipService()

	pk := domain.UserPubKey(params.Pubkey)

	pks, err := relationShipService.GetFollowersPubKeys(pk)
	if err != nil {
		return err
	}

	return c.JSON(
		http.StatusOK,
		ToPubKeysResponse(pks),
	)
}
