//go:build wireinject && !mock

package injection

import (
	"github.com/google/wire"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/repository"
	"github.com/uakihir0/nostr-rest/server/service"
)

var bindSet = wire.NewSet(
	service.NewUserService,
	service.NewPostService,
	service.NewRelationShipService,
	repository.NewRelayUserRepository,
	repository.NewRelayPostRepository,
	repository.NewRelayRelationShipRepository,
	wire.Bind(
		new(domain.UserRepository),
		new(*repository.RelayUserRepository),
	),
	wire.Bind(
		new(domain.PostRepository),
		new(*repository.RelayPostRepository),
	),
	wire.Bind(
		new(domain.RelationShipRepository),
		new(*repository.RelayRelationShipRepository),
	),
)
