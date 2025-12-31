package main

import (
	"fmt"
	"unsafe"
)

type Vector[T any] struct {
	len uint64
	cap uint64
	ptr unsafe.Pointer
}

func NewVect[T any](cap uint64) *Vector[T] {
	v := &Vector[T]{len: 0, cap: cap, ptr: nil}
	if cap > 0 {
		v.grow(cap)
	}
	return v
}

func (v *Vector[T]) grow(newCap uint64) {
	var zero T
	elementSize := unsafe.Sizeof(zero)

	newMem := make([]T, newCap)
	newPtr := unsafe.Pointer(&newMem[0])

	if v.ptr != nil {
		for i := uint64(0); i < v.len; i++ {

			src := unsafe.Pointer(uintptr(v.ptr) + (uintptr(i) * elementSize))

			dst := unsafe.Pointer(uintptr(newPtr) + (uintptr(i) * elementSize))

			*(*T)(dst) = *(*T)(src)
		}
	}

	v.ptr = newPtr
	v.cap = newCap

	fmt.Printf("Grown! New Cap: %d, Addr: %v\n", v.cap, v.ptr)
}

func (v *Vector[T]) PushBack(value T) {
	if v.len == v.cap {
		newcap := v.cap * 2
		if newcap == 0 {
			newcap = 1
		}
		v.grow(newcap)
	}
	elementSize := unsafe.Sizeof(value)
	offset := uintptr(v.len) * elementSize

	targetAddr := unsafe.Pointer(uintptr(v.ptr) + offset)

	*(*T)(targetAddr) = value

	v.len++
}

func (v *Vector[T]) Size() uint64 {
	return v.len
}

func (v *Vector[T]) Capacity() uint64 {
	return v.cap
}

func (v *Vector[T]) PopBack() {
	if v.len == 0 {
		return
	}
	v.len--
	var zero T
	elementSize := unsafe.Sizeof(zero)
	offset := uintptr(v.len) * elementSize
	targetAddr := unsafe.Pointer(uintptr(v.ptr) + offset)
	*(*T)(targetAddr) = zero
}

func (v *Vector[T]) At(index uint64) T {
	if index >= v.len {
		panic("Index out of bounds")
	}
	var zero T
	elementSize := unsafe.Sizeof(zero)

	targetAddr := unsafe.Pointer(uintptr(v.ptr) + (uintptr(index) * elementSize))
	return *(*T)(targetAddr)
}
