package domain

// UserPubKey User's public key
type UserPubKey string

// UserSecretKey User's secret key
type UserSecretKey string

type User struct {
	PubKey      UserPubKey
	Name        string
	DisplayName string
	About       string
	Picture     string
	Banner      string
	Website     string
}

type UserRepository interface {

	// GetUsers
	GetUsers([]UserPubKey) ([]*User, error)
}
