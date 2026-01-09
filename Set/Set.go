package Set

import (
	"sync"
)

type Set[T comparable] interface {
	Add(T)
	Contains(T) bool
	Remove(T)
	Clear()
}

type mapset[T comparable] struct {
	mu sync.RWMutex
	s  map[T]struct{}
}

func NewSet[T comparable]() *mapset[T] {
	return &mapset[T]{s: make(map[T]struct{})}
}

func (ms *mapset[T]) Add(element T) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.s[element] = struct{}{}
}

func (ms *mapset[T]) Remove(element T) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	delete(ms.s, element)
}

func (ms *mapset[T]) Contains(element T) bool {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	_, found := ms.s[element]
	return found
}

func (ms *mapset[T]) Clear() {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	for k := range ms.s {
		delete(ms.s, k)
	}

}
