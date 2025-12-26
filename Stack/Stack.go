package stack

import (
	"errors"
	"sync"
)

type Stack interface {
	Push(val any)
	Pop() (any, error)
}

type threadSafeStack struct {
	data []any
	mu   sync.RWMutex
}

func (s *threadSafeStack) Push(val any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, val)
}

func (s *threadSafeStack) Pop() (any, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.data) == 0 {
		return nil, errors.New("stack is empty")
	}

	lastIndex := len(s.data) - 1
	element := s.data[lastIndex]
	s.data = s.data[:lastIndex]

	return element, nil
}

func NewStack() Stack {
	return &threadSafeStack{}
}
