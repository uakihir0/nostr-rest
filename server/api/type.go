package api

import (
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/openapi"
)

func ToUser(user *domain.User) *openapi.User {
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
		Count: len(users),
		List: lo.Map(users,
			func(u *domain.User, _ int) openapi.User {
				return *ToUser(u)
			},
		),
	}
}

func ToPubKeysResponse(pks []domain.UserPubKey) *openapi.PubKeysResponse {
	return &openapi.PubKeysResponse{
		Pubkeys: lo.Map(pks,
			func(i domain.UserPubKey, _ int) string {
				return string(i)
			},
		),
	}
}

func ToTimeline(posts []*domain.Post, users []*domain.User) *openapi.PostsResponse {

	userMap := make(map[string]*domain.User)
	for _, user := range users {
		pk := string(user.PubKey)
		userMap[pk] = user
	}

	postResponses := make([]openapi.Post, 0)
	for _, post := range posts {
		pk := string(post.UserPubKey)
		user, ok := userMap[pk]

		if ok {
			postResponses = append(
				postResponses,
				openapi.Post{
					Content: post.Content,
					User:    *ToUser(user),
				},
			)
		}
	}

	return &openapi.PostsResponse{
		Count: len(postResponses),
		List:  postResponses,
	}
}
