package domain

import (
	"time"
)

type PostID string

type Post struct {
	ID         PostID
	UserPubKey UserPubKey
	Content    string
	CreatedAt  time.Time

	RootPostID           *PostID
	ReplyPostID          *PostID
	MentionedUserPubKeys []UserPubKey
}

type PostWithUser struct {
	Post
	User
}

type PostRepository interface {

	// SendPost
	SendPost(UserPubKey, UserSecretKey, string) error

	// GetPost
	GetPost(id PostID) (*Post, error)

	// GetPosts
	GetPosts(
		pks []UserPubKey,
		options PagingOptions,
	) ([]Post, error)

	// GetPublicPosts
	GetPublicPosts(
		options PagingOptions,
	) ([]Post, error)

	// GetUserLatestPosts
	GetUserLatestPosts(
		pk UserPubKey,
	) ([]Post, error)

	// GetReplies
	GetReplies(PostID) ([]Post, error)
}
