package minstack

import (
	"cmp"
	"fmt"
	"sync"
)

type MinStack[T cmp.Ordered] interface {
	Push(value T)
	Pop() (T, error)
	Top() T
	Min() T
	String() string
}

type minstack[T cmp.Ordered] struct {
	mu      sync.RWMutex
	mindata []T
	data    []T
}

func NewMinStack[T cmp.Ordered]() *minstack[T] {
	return &minstack[T]{mindata: make([]T, 0), data: make([]T, 0)}
}

func (st *minstack[T]) Push(value T) {
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

func (st *minstack[T]) Pop() (T, error) {
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

func (st *minstack[T]) Top() T {
	st.mu.RLock()
	defer st.mu.RUnlock()
	if len(st.data) == 0 {
		var zero T
		return zero
	}
	return st.data[len(st.data)-1]
}

func (st *minstack[T]) Min() T {
	st.mu.RLock()
	defer st.mu.RUnlock()
	if len(st.mindata) == 0 {
		var zero T
		return zero
	}
	return st.mindata[len(st.mindata)-1]
}

func (st *minstack[T]) String() string {
	return fmt.Sprintln("MinStack = ", st.mindata, " Stack = ", st.data)

}
