package domain

import "time"

type RepostID string

type Repost struct {
	ID         RepostID
	UserPubKey UserPubKey
	Content    string
	CreatedAt  time.Time
}

type RepostRepository interface {

	// GetReposts
	GetReposts(PostID) ([]Repost, error)
}
