package repository

import (
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/nbd-wtf/go-nostr"
)

type RelayEventsCache struct {
	Events    []*nostr.Event
	ExpiredAt int64
}

type StringCacheMap = *lru.Cache[string, *RelayEventsCache]

func NewStringCacheMap(size int) StringCacheMap {
	c, err := lru.New[string, *RelayEventsCache](size)
	if err != nil {
		panic("Error on NewUserPubKeyCacheMap Init")
	}
	return c
}

// GetEventsFromString
func GetEventsFromString(
	c StringCacheMap,
	key string,
) []*nostr.Event {
	value, ok := c.Get(key)
	if !ok {
		return nil
	}

	now := time.Now().Unix()
	if value.ExpiredAt <= now {
		c.Remove(key)
		return nil
	}
	return value.Events
}

// SetEventsFromString
func SetEventsFromString(
	c StringCacheMap,
	key string,
	v []*nostr.Event,
) {
	c.Add(
		key,
		&RelayEventsCache{
			Events:    v,
			ExpiredAt: time.Now().Unix() + 60,
		},
	)
}

// ManageEventsFromString
func ManageEventsFromString(
	c StringCacheMap,
	key string,
	supplier func() []*nostr.Event,
) []*nostr.Event {

	events := GetEventsFromString(c, key)
	if events == nil {
		events = supplier()
		SetEventsFromString(c, key, events)
	}
	return events
}
