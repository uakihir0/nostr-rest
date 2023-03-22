package domain

type RelationShipRepository interface {

	// GetFollowings
	GetFollowings(UserPubKey) ([]UserPubKey, error)

	// GetFollowers
	GetFollowers(UserPubKey) ([]UserPubKey, error)
}
