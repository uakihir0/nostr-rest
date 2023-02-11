package main

import (
	"github.com/uakihir0/nostr-rest/server"
	"github.com/uakihir0/nostr-rest/server/repository"
)

func main() {

	repository.ConnectAllRelayServers(
		[]string{
			"wss://relay.damus.io",
			"wss://relay.snort.social",
		},
	)

	server.Serve()
}
