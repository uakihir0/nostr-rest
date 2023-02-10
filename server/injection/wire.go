//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/uakihir0/nostr-rest/server/domain"
)

func UserRepository() *domain.UserRepository {
	wire.Build(bindSet)
	return nil
}