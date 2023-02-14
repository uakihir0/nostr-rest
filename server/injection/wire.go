//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/uakihir0/nostr-rest/server/service"
)

func UserService() *service.UserService {
	wire.Build(bindSet)
	return nil
}

func RelationShipService() *service.RelationShipService {
	wire.Build(bindSet)
	return nil
}