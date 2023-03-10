package repository

import (
	"context"
	"fmt"
	"github.com/nbd-wtf/go-nostr"
	"github.com/samber/lo"
	"sync"
	"sync/atomic"
	"time"
)

type RelayConnection struct {
	URL      string
	Relay    *nostr.Relay
	Context  *context.Context
	Cancel   *context.CancelFunc
	WaitSpan time.Duration
}

var connections = map[string]*RelayConnection{}

// GetConnections
// Obtain relay server connection information
func GetConnections() []*RelayConnection {
	return lo.Filter(lo.Values(connections),
		func(e *RelayConnection, _ int) bool {
			// Get valid connections
			return e.Relay != nil
		})
}

// StartAllRelayServers
// Connect to all relay servers.
func StartAllRelayServers(urls []string) {
	for _, url := range urls {
		StartRelayServer(url)
	}
}

// StartRelayServer
// Connect to relay server.
func StartRelayServer(url string) {

	c := &RelayConnection{URL: url, WaitSpan: 1}
	connections[url] = c

	go func() {
		for {
			err := ConnectRelay(c)
			if err != nil {
				PrintConnectionError(c, err)
				SleepTimeSpan(c)
				continue
			}

			// Print connection succeeded
			fmt.Printf("> connection success: %s\n", c.URL)

			select {
			// Waiting for connection error
			case err = <-c.Relay.ConnectionError:
				ClearRelay(c)

				PrintConnectionError(c, err)
				SleepTimeSpan(c)
			}
		}
	}()
}

func ConnectRelay(c *RelayConnection) error {

	ctx, cancel := context.WithCancel(context.Background())
	relay, err := nostr.RelayConnect(ctx, c.URL)
	if err != nil {
		cancel()
		return err
	}

	c.Relay = relay
	c.Context = &ctx
	c.Cancel = &cancel
	return nil
}

func ClearRelay(c *RelayConnection) {
	c.Relay = nil
	c.Context = nil
	c.Cancel = nil
}

func SleepTimeSpan(c *RelayConnection) {

	// Insert Timespan to wait
	time.Sleep(c.WaitSpan * time.Second)
	c.WaitSpan = c.WaitSpan * 2
	if c.WaitSpan > 60 {
		c.WaitSpan = 60
	}
}

func PrintConnectionError(c *RelayConnection, err error) {
	fmt.Printf("> connection error: %s\n", c.URL)
	fmt.Printf(">> error: %s\n", err)
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
	stop := make(chan interface{})

	events := make([]*nostr.Event, 0)
	dones := make([]chan interface{}, 0)

	// Channel close mutex
	var mu sync.Mutex
	var isStopped = false
	var timeout = 1 * time.Second

	afterChannelClose := func() {
		// Close channel if all subscriptions closed
		atomic.AddInt32(&channelCounter, 1)

		mu.Lock()
		if channelCounter == int32(len(cs)) {
			if !isStopped {
				isStopped = true
				close(stop)
			}
		}
		mu.Unlock()
	}

	for _, c := range cs {
		// Timeout occurs if acquisition is not possible
		ctx, cancel := context.WithTimeout(ctx, timeout)
		sub := c.Relay.Subscribe(ctx, filters)

		done := make(chan interface{})
		dones = append(dones, done)

		go func(done <-chan interface{}) {

			defer func() {
				afterChannelClose()
				sub.Unsub()
				cancel()
			}()

			for {
				select {
				// Termination process first
				case <-sub.EndOfStoredEvents:
					return
				case <-ctx.Done():
					return
				case <-done:
					return

				case ev := <-sub.Events:
					channel <- ev
					continue
				}
			}
		}(done)
	}

loop:
	for {
		select {
		case <-stop:
			break loop

		case ev := <-channel:
			events = append(events, ev)

			mu.Lock()
			if expectedEventCount > 0 {
				// Early termination (got expected data)
				if len(events) >= expectedEventCount {
					if !isStopped {
						isStopped = true
						close(stop)

						// Send done signal to connections.
						for _, done := range dones {
							close(done)
						}
					}
					break loop
				}
			}
			mu.Unlock()
			continue
		}
	}

	return events
}
