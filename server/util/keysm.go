package util

import (
	"context"
	"github.com/samber/lo"
	"golang.org/x/sync/semaphore"
	"sync"
)

type LimitMap struct {
	refs   map[string]*int
	limits map[string]*semaphore.Weighted
	lock   sync.Mutex
	max    int
}

func NewLimitMap(max int) *LimitMap {
	return &LimitMap{
		max:    max,
		refs:   make(map[string]*int),
		limits: make(map[string]*semaphore.Weighted),
	}
}

// Acquire
func (m *LimitMap) Acquire(ctx context.Context, key string) {
	m.lock.Lock()
	l, ok := m.limits[key]
	if !ok {
		l = semaphore.NewWeighted(int64(m.max))
		m.limits[key] = l
	}
	c, ok := m.refs[key]
	if !ok {
		c = lo.ToPtr(0)
	}
	c = lo.ToPtr(*c + 1)
	m.refs[key] = c
	m.lock.Unlock()

	err := l.Acquire(ctx, 1)
	if err != nil {
		panic(err)
	}
}

// Release
func (m *LimitMap) Release(key string) {
	m.lock.Lock()
	l, ok := m.limits[key]
	if !ok {
		panic("LimitMap: key not in map. Possible reason: Release without Acquire.")
	}
	c, ok := m.refs[key]
	if !ok {
		panic("LimitMap: key not in map. Possible reason: Release without Acquire.")
	}

	c = lo.ToPtr(*c - 1)
	m.refs[key] = c

	if *c <= 0 {
		delete(m.limits, key)
		delete(m.refs, key)
	}
	m.lock.Unlock()

	l.Release(1)
}
