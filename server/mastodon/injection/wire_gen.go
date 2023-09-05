// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package minjection

import (
	"github.com/uakihir0/nostr-rest/server/mastodon/service"
	"github.com/uakihir0/nostr-rest/server/repository"
)

// Injectors from wire.go:

func AccountService() *mservice.AccountService {
	relayUserRepository := repository.NewRelayUserRepository()
	relayPostRepository := repository.NewRelayPostRepository()
	relayRepostRepository := repository.NewRelayRepostRepository()
	relayReactionRepository := repository.NewRelayReactionRepository()
	relayRelationShipRepository := repository.NewRelayRelationShipRepository()
	typeService := mservice.NewTypeService(relayUserRepository, relayPostRepository, relayRepostRepository, relayReactionRepository, relayRelationShipRepository)
	accountService := mservice.NewAccountService(typeService, relayUserRepository, relayPostRepository, relayRelationShipRepository)
	return accountService
}

func StatusService() *mservice.StatusService {
	relayUserRepository := repository.NewRelayUserRepository()
	relayPostRepository := repository.NewRelayPostRepository()
	relayRepostRepository := repository.NewRelayRepostRepository()
	relayReactionRepository := repository.NewRelayReactionRepository()
	relayRelationShipRepository := repository.NewRelayRelationShipRepository()
	typeService := mservice.NewTypeService(relayUserRepository, relayPostRepository, relayRepostRepository, relayReactionRepository, relayRelationShipRepository)
	statusService := mservice.NewStatusService(typeService, relayUserRepository, relayPostRepository, relayRelationShipRepository)
	return statusService
}

func TimelineService() *mservice.TimelineService {
	relayUserRepository := repository.NewRelayUserRepository()
	relayPostRepository := repository.NewRelayPostRepository()
	relayRepostRepository := repository.NewRelayRepostRepository()
	relayReactionRepository := repository.NewRelayReactionRepository()
	relayRelationShipRepository := repository.NewRelayRelationShipRepository()
	typeService := mservice.NewTypeService(relayUserRepository, relayPostRepository, relayRepostRepository, relayReactionRepository, relayRelationShipRepository)
	timelineService := mservice.NewTimelineService(typeService, relayUserRepository, relayPostRepository, relayRelationShipRepository)
	return timelineService
}
