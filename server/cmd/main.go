package main

import (
	"fmt"
	"github.com/uakihir0/nostr-rest/server/domain"
	"github.com/uakihir0/nostr-rest/server/repository"
)

func main() {

	repository.ConnectAllRelayServers(
		[]string{
			"wss://relay.damus.io",
			"wss://relay.snort.social",
		},
	)

	userRepository := repository.NewRelayUserRepository()

	users, err := userRepository.GetUsers([]domain.UserPubKey{
		"776ea4437354381f14a720be3c476937dce7257ed1073e54a192dbc99f3b7ecc",
	})
	if err != nil {
		return
	}

	for _, user := range users {
		fmt.Println(user.PubKey)
		fmt.Println(user.Name)
		fmt.Println(user.DisplayName)
		fmt.Println(user.About)
	}
}
