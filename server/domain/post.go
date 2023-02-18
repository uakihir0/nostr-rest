package domain

import "time"

type PostID string

type Post struct {
	ID         PostID
	UserPubKey UserPubKey
	Content    string
	CreatedAt  time.Time
}

type PostWithUser struct {
	Post
	User
}

type PostRepository interface {

	// GetPosts
	GetPosts([]UserPubKey, int, *time.Time, *time.Time) ([]*Post, error)
}
