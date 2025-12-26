package stack

import (
	"errors"
	"sync"
)

type Stack[T any] interface {
	Push(val T)
	Pop() (T, error)
}

type threadSafeStack[T any] struct {
	data []T
	mu   sync.RWMutex
}

func (s *threadSafeStack[T]) Push(val T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, val)
}

func (s *threadSafeStack[T]) Pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var zero T
	if len(s.data) == 0 {
		return zero, errors.New("stack is empty")
	}

	lastIndex := len(s.data) - 1
	element := s.data[lastIndex]
	s.data = s.data[:lastIndex]

	return element, nil
}

func (s *threadSafeStack[T]) Empty() bool {
	return len(s.data) == 0
}

func NewStack[T any]() Stack[T] {
	return &threadSafeStack[T]{}
}
