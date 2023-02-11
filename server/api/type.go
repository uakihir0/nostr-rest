package api

import (
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/openapi"
)

func ToUserResponse(user *domain.User) *openapi.User {

	return &openapi.User{
		Pubkey:      string(user.PubKey),
		Name:        lo.ToPtr(user.Name),
		DisplayName: lo.ToPtr(user.DisplayName),
		About:       lo.ToPtr(user.About),
		Picture:     lo.ToPtr(user.Picture),
		Banner:      lo.ToPtr(user.Banner),
		Website:     lo.ToPtr(user.Website),
	}
}

func ToUsersResponse(users []*domain.User) *openapi.UsersResponse {

	return &openapi.UsersResponse{
		List: lo.Map(users,
			func(u *domain.User, _ int) openapi.User {
				return *ToUserResponse(u)
			}),
	}
}
