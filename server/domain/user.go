package domain

import (
	"errors"
	"github.com/nbd-wtf/go-nostr/nip19"
	"strings"
)

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
	Lud06       string
	CreatedAt   int64
}

func ToUserPubKey(uid string) (*UserPubKey, error) {

	// uid is starting npub format
	if strings.HasPrefix(uid, "npub") {
		_, decoded, err := nip19.Decode(uid)
		if err != nil {
			return nil, errors.New("illegal public key")
		}
		pk := UserPubKey(decoded.(string))
		return &pk, nil
	}

	// uid is hash format
	pk := UserPubKey(uid)
	return &pk, nil
}

type UserRepository interface {

	// GetUsers
	GetUsers([]UserPubKey) ([]*User, error)
}
