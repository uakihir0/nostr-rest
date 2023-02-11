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
func (h *Handler) GetV1Users(c echo.Context, params openapi.GetV1UsersParams) error {
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
func (h *Handler) PostV1Users(c echo.Context) error {
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
