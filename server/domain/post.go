package domain

import "time"

type Post struct {
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
	GetPosts([]UserPubKey) ([]*Post, error)
}
