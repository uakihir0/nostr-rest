package domain

// UserPubKey ユーザーの公開鍵
type UserPubKey string

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
