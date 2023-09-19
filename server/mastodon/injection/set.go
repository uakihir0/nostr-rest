package minjection

import (
	"github.com/google/wire"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/mastodon/service"
	"github.com/uakihir0/nostr-rest/server/repository"
)

var bindSet = wire.NewSet(
	mservice.NewTypeService,
	mservice.NewAccountService,
	mservice.NewStatusService,
	mservice.NewTimelineService,
	repository.NewRelayUserRepository,
	repository.NewRelayPostRepository,
	repository.NewRelayRepostRepository,
	repository.NewRelayReactionRepository,
	repository.NewRelayTimelineRepository,
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
		new(domain.RepostRepository),
		new(*repository.RelayRepostRepository),
	),
	wire.Bind(
		new(domain.ReactionRepository),
		new(*repository.RelayReactionRepository),
	),
	wire.Bind(
		new(domain.TimelineRepository),
		new(*repository.RelayTimelineRepository),
	),
	wire.Bind(
		new(domain.RelationShipRepository),
		new(*repository.RelayRelationShipRepository),
	),
)
