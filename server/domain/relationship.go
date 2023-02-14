package domain

type RelationShipRepository interface {

	// GetFollowings
	GetFollowings(UserPubKey) ([]UserPubKey, error)
}
