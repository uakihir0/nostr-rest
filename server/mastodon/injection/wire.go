//go:build wireinject
// +build wireinject

package minjection

import (
	"github.com/google/wire"
	"github.com/uakihir0/nostr-rest/server/mastodon/service"
)

func AccountService() *mservice.AccountService {
	wire.Build(bindSet)
	return nil
}

func StatusService() *mservice.StatusService {
	wire.Build(bindSet)
	return nil
}

func TimelineService() *mservice.TimelineService {
	wire.Build(bindSet)
	return nil
}
