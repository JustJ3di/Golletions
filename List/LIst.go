package list

import (
	"fmt"
	"strings"
)

type node[T any] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

func (it *node[T]) Next() *node[T] {
	return it.next
}

func (it *node[T]) Prev() *node[T] {
	return it.prev
}

type List[T comparable] struct {
	len  int
	head *node[T]
	tail *node[T]
}

func New[T comparable]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Push(value T) {
	newNode := &node[T]{value: value, next: l.head, prev: nil}
	if l.head != nil {
		l.head.prev = newNode
	}
	l.head = newNode
	if l.tail == nil {
		l.tail = newNode
	}
	l.len++
}

func (l *List[T]) Mpush(value ...T) {
	for _, v := range value {
		l.Push(v)
	}
}

func (l *List[T]) Mpop(npop int) []T {
	rt := make([]T, npop)
	for i := range npop {
		var fl bool
		rt[i], fl = l.Pop()
		if fl == false {
			return nil
		}
	}
	return rt
}

func (l *List[T]) Pop() (T, bool) {
	var zero T
	if l.head == nil {
		return zero, false
	}
	val := l.head.value
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	l.len--
	return val, true
}

func (l *List[T]) Remove(value T) {
	curr := l.head
	for curr != nil {
		if curr.value == value {
			if curr.prev != nil {
				curr.prev.next = curr.next
			} else {
				l.head = curr.next
			}
			if curr.next != nil {
				curr.next.prev = curr.prev
			} else {
				l.tail = curr.prev
			}
			l.len--
		}
		curr = curr.next
	}
}

func (l *List[T]) Reverse() {
	if l.head == nil {
		return
	}
	curr := l.head
	for curr != nil {
		prev := curr.prev
		curr.prev = curr.next
		curr.next = prev
		curr = curr.prev
	}
	temp := l.head
	l.head = l.tail
	l.tail = temp
}

func (l *List[T]) Len() int {
	return l.len
}

func (l *List[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for t := l.head; t != nil; t = t.next {
		sb.WriteString(fmt.Sprintf("%v", t.value))
		if t.next != nil {
			sb.WriteString(" ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}
