package main

import (
	"github.com/uakihir0/nostr-rest/server"
	"github.com/uakihir0/nostr-rest/server/repository"
)

func main() {

	repository.ConnectAllRelayServers(
		[]string{
			"wss://eden.nostr.land",
			"wss://relay.damus.io",
			"wss://nos.lol",
			"wss://brb.io",
			"wss://relay.snort.social",
		},
	)

	server.Serve()
}
