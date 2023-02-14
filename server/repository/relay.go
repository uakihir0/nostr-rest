package repository

import (
	"context"
	"github.com/nbd-wtf/go-nostr"
	"github.com/samber/lo"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type RelayConnection struct {
	URL      string
	Relay    *nostr.Relay
	Context  *context.Context
	Cancel   *context.CancelFunc
	WaitSpan int
}

var connections = map[string]*RelayConnection{}

// GetConnections
// Obtain relay server connection information
func GetConnections() []*RelayConnection {
	return lo.Filter(lo.Values(connections),
		func(e *RelayConnection, _ int) bool {
			// Get valid connections with no errors
			return e.Relay != nil && e.Relay.ConnectionError != nil
		})
}

// ConnectAllRelayServers
// Connect to all relay servers.
func ConnectAllRelayServers(urls []string) {
	for _, url := range urls {
		ConnectRelayServer(url)
	}
}

// ConnectRelayServer
// Connect to relay server.
func ConnectRelayServer(url string) {

	ctx, cancel := context.WithCancel(context.Background())
	relay, err := nostr.RelayConnect(ctx, url)

	if err != nil {
		// If connection is not made, reconnect later.
		connection := &RelayConnection{URL: url, WaitSpan: 1}
		connections[url] = connection
		cancel()

	} else {
		connection := &RelayConnection{
			URL:      url,
			Relay:    relay,
			Context:  &ctx,
			Cancel:   &cancel,
			WaitSpan: 1,
		}
		connections[url] = connection
	}
}

// QuerySyncAll
// Query all relay servers and retrieve results synchronously
func QuerySyncAll(
	ctx context.Context,
	filters nostr.Filters,
) []*nostr.Event {
	return QuerySyncAllWithGuard(ctx, filters, -1)
}

// QuerySyncAllWithGuard
// Query all relay servers and retrieve results synchronously
func QuerySyncAllWithGuard(
	ctx context.Context,
	filters nostr.Filters,
	expectedEventCount int,
) []*nostr.Event {

	cs := GetConnections()

	var channelCounter int32 = 0
	channel := make(chan *nostr.Event)
	events := make([]*nostr.Event, 0)

	// Channel close mutex
	var mu sync.Mutex
	var isClosed = false

	checkChannelClose := func() {
		// Close channel if all subscriptions closed
		atomic.AddInt32(&channelCounter, 1)

		mu.Lock()
		if !isClosed && channelCounter == int32(len(cs)) {
			close(channel)
			isClosed = true
		}
		mu.Unlock()
	}

	for _, c := range cs {
		// Timeout occurs if acquisition is not possible
		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		sub := c.Relay.Subscribe(ctx, filters)

		go func() {
			for {
				select {
				case ev := <-sub.Events:
					// Send event to channel
					mu.Lock()
					if !isClosed {
						channel <- ev
					}
					mu.Unlock()

				case <-sub.EndOfStoredEvents:
					// Normal termination process
					checkChannelClose()
					sub.Unsub()
					cancel()
					return

				case <-ctx.Done():
					checkChannelClose()
					sub.Unsub()
					return
				}
			}
		}()
	}

	// Consolidate events
	for ev := range channel {
		events = append(events, ev)

		mu.Lock()
		if !isClosed && expectedEventCount > 0 {
			// Early termination (got expected data)
			if len(events) >= expectedEventCount {
				close(channel)
				isClosed = true
			}
		}
		mu.Unlock()
	}

	// Sort in descending order of `CreatedAt`
	sort.Slice(events, func(i, j int) bool {
		return events[i].CreatedAt.Unix() < events[j].CreatedAt.Unix()
	})

	return events
}
