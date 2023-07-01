package api

import (
	"github.com/labstack/echo/v4"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/injection"
	"github.com/uakihir0/nostr-rest/server/openapi"
)

// PostV1Posts
func (h *SimpleHandler) PostV1Posts(c echo.Context) error {
	postService := injection.PostService()

	request := new(openapi.PostCommentRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	err := postService.SendPost(
		domain.UserPubKey(request.Keys.Public),
		domain.UserSecretKey(request.Keys.Secret),
		request.Text,
	)
	if err != nil {
		return err
	}
	return nil
}
