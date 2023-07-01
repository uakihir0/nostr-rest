package main

import (
	"github.com/uakihir0/nostr-rest/server"
	"github.com/uakihir0/nostr-rest/server/repository"
)

func main() {

	repository.StartAllRelayServers(
		[]string{
			"wss://relay.snort.social",
			"wss://eden.nostr.land",
			"wss://relay.damus.io",
			"wss://nos.lol",
		},
	)

	server.Serve()
}
