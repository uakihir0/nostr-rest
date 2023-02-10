//go:build wireinject && !mock

package injection

import (
	"github.com/google/wire"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/repository"
)

var bindSet = wire.NewSet(
	repository.NewRelayUserRepository,
	wire.Bind(
		new(domain.UserRepository),
		new(*repository.RelayUserRepository),
	),
)
