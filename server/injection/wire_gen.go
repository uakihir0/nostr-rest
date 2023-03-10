// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/uakihir0/nostr-rest/server/repository"
	"github.com/uakihir0/nostr-rest/server/service"
)

// Injectors from wire.go:

func UserService() *service.UserService {
	relayUserRepository := repository.NewRelayUserRepository()
	userService := service.NewUserService(relayUserRepository)
	return userService
}

func PostService() *service.PostService {
	relayPostRepository := repository.NewRelayPostRepository()
	postService := service.NewPostService(relayPostRepository)
	return postService
}

func RelationShipService() *service.RelationShipService {
	relayRelationShipRepository := repository.NewRelayRelationShipRepository()
	relationShipService := service.NewRelationShipService(relayRelationShipRepository)
	return relationShipService
}
