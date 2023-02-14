package util

import "sync"

type Lock[T any] struct {
	once     sync.Once
	instance *T
}

func (l *Lock[T]) Once(f func() *T) *T {
	l.once.Do(func() {
		l.instance = f()
	})
	return l.instance
}
