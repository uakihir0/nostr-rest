package minjection

import (
	"github.com/google/wire"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/mastodon/service"
	"github.com/uakihir0/nostr-rest/server/repository"
)

var bindSet = wire.NewSet(
	mservice.NewAccountService,
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
