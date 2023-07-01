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
	relayRelationShipRepository := repository.NewRelayRelationShipRepository()
	accountService := mservice.NewAccountService(relayUserRepository, relayPostRepository, relayRelationShipRepository)
	return accountService
}
