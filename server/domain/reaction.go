package domain

import "time"

type ReactionID string

type Reaction struct {
	ID         ReactionID
	UserPubKey UserPubKey
	Content    string
	CreatedAt  time.Time

	ToPostID     PostID
	ToUserPubKey UserPubKey
}

type ReactionRepository interface {

	// GetReactions
	GetReactions(PostID) ([]Reaction, error)
}
