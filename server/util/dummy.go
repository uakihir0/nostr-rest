package util

import (
	"github.com/samber/lo"
	"github.com/uakihir0/nostr-rest/server/domain"
)

func GetNoDataUser(pk domain.UserPubKey) *domain.User {
	name := "NoData:" + lo.Substring(string(pk), 0, 10)

	return &domain.User{
		PubKey:      pk,
		Name:        name,
		DisplayName: name,
	}
}
