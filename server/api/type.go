package api

import (
	"github.com/uakihir0/nostr-rest/server/util"
	"time"

	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/openapi"
)

const (
	TimeLayout = "2006-01-02 15:04:05"
)

func ToUser(user domain.User) *openapi.User {
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

func ToUsersResponse(users []domain.User) *openapi.UsersResponse {
	return &openapi.UsersResponse{
		Count: len(users),
		List: lo.Map(users,
			func(u domain.User, _ int) openapi.User {
				return *ToUser(u)
			},
		),
	}
}

func ToPubKeysResponse(pks []domain.UserPubKey) *openapi.PubKeysResponse {
	return &openapi.PubKeysResponse{
		Count: len(pks),
		Pubkeys: lo.Map(pks,
			func(i domain.UserPubKey, _ int) string {
				return string(i)
			},
		),
	}
}

func ToTimeline(
	pks []domain.UserPubKey,
	posts []domain.Post,
	users []domain.User,
) *openapi.UsersTimelineResponse {

	userMap := make(map[string]domain.User)
	for _, user := range users {
		pk := string(user.PubKey)
		userMap[pk] = user
	}

	postSlice := make([]openapi.Post, 0)
	for _, post := range posts {
		pk := string(post.UserPubKey)
		user, ok := userMap[pk]

		if !ok {
			// User information cannot be obtained
			user = util.GetNoDataUser(post.UserPubKey)
		}

		postSlice = append(postSlice, openapi.Post{
			Id:        string(post.ID),
			CreatedAt: post.CreatedAt.Format(TimeLayout),
			Content:   post.Content,
			User:      *ToUser(user),
		})
	}

	postsResponse := openapi.Posts{
		Count: len(postSlice),
		List:  postSlice,
	}

	var pagingResponse *openapi.Paging
	if len(postSlice) > 0 {

		lastPost, err := lo.Last(posts)
		if err != nil {
			return nil
		}

		pagingResponse = &openapi.Paging{
			FutureSinceTime: posts[0].CreatedAt.Add(+time.Second).Format(TimeLayout),
			PastUntileTime:  lastPost.CreatedAt.Format(TimeLayout),
		}
	}

	pksResponse := lo.Map(pks,
		func(pk domain.UserPubKey, _ int) string {
			return string(pk)
		})

	return &openapi.UsersTimelineResponse{
		Posts:   postsResponse,
		Paging:  pagingResponse,
		Pubkeys: pksResponse,
	}
}
