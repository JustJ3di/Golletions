package minstack

import (
	"cmp"
	"fmt"
	"sync"
)

type MinStack[T cmp.Ordered] struct {
	mu      sync.RWMutex
	mindata []T
	data    []T
}

func New[T cmp.Ordered]() *MinStack[T] {
	return &MinStack[T]{mindata: make([]T, 0), data: make([]T, 0)}
}

func (st *MinStack[T]) Push(value T) {
	st.mu.Lock()
	defer st.mu.Unlock()
	st.data = append(st.data, value)
	if len(st.mindata) == 0 {
		st.mindata = append(st.mindata, value)
	}
	lstidx := len(st.mindata) - 1
	if st.mindata[lstidx] > value {
		st.mindata = append(st.mindata, value)
	}
}

func (st *MinStack[T]) Pop() (T, error) {
	st.mu.Lock()
	defer st.mu.Unlock()
	var ret T
	if len(st.data) == 0 {
		return ret, fmt.Errorf("Stack empty")
	} else {
		ret = st.data[len(st.data)-1]
		st.data = st.data[:len(st.data)-1]

		if ret == st.mindata[len(st.mindata)-1] {
			st.mindata = st.mindata[:len(st.mindata)-1]
		}
	}
	return ret, nil
}

func (st *MinStack[T]) Top() T {
	st.mu.RLock()
	defer st.mu.RUnlock()
	return st.data[len(st.data)-1]
}

func (st *MinStack[T]) Min() T {
	st.mu.RLock()
	defer st.mu.RUnlock()
	return st.mindata[len(st.mindata)-1]
}

func (st *MinStack[T]) String() string {
	return fmt.Sprintln("MinStack = ", st.mindata, " Stack = ", st.data)
}
