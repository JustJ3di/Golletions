# Go Unsafe Vector 🚀

An experimental, generic implementation of a dynamic array (**Vector**) in Go, built to replicate the internal memory management behavior of C++ `std::vector`.

This project demonstrates **manual memory management**, **pointer arithmetic**, and the usage of Go's `unsafe` package to bypass standard slice mechanics.

> **⚠️ DISCLAIMER**: This project is intended for **educational purposes**.  
> While functional, this code bypasses Go's memory safety checks. In idiomatic Go, you should almost always use standard slices (`[]T`). Use this library if you want to understand how vectors work "under the hood" or need to interface with raw memory manually.

## ✨ Features

* **Generics:** Fully compatible with Go 1.18+ (`Vector[T]`).
* **Manual Memory Layout:** Uses `unsafe.Pointer` to manage a contiguous block of memory, bypassing Go slice headers.
* **Dynamic Growth:** Automatically doubles capacity when full (Geometric Growth strategy).
* **GC-Friendly PopBack:** Implements specific "zeroing out" logic to prevent memory leaks (loitering objects) within the Garbage Collector.
* **Direct Pointer Arithmetic:** Uses `uintptr` for element access.

## 📦 Installation

Since this is a lightweight data structure, you can simply drop the implementation file into your project.

## 🚀 Usage Example

```go
package main

import (
	"fmt"
	"yourproject/vector" // Import your package path
)

func main() {
	// 1. Create a new Vector (similar to std::vector<int> v)
	// Initial capacity: 2
	v := vector.NewVect[int](2)

	// 2. Push elements (triggers automatic allocation and growth)
	v.PushBack(10)
	v.PushBack(20)
	v.PushBack(30) // Capacity doubles here automatically

	// 3. Check metadata
	fmt.Printf("Size: %d, Capacity: %d\n", v.Size(), v.Capacity())

	// 4. Access elements (assuming an At() method is implemented)
	// fmt.Println(v.At(1)) // Prints 20

	// 5. Pop elements safely
	v.PopBack() // Removes 30 and clears memory reference for the GC
}
```

### 🧠 Under the Hood
Memory Layout

Unlike a Go slice (which is a "view" structure pointing to an array), this Vector owns its memory pointer directly, mimicking the C layout.
Plaintext

Vector Struct       Heap Memory (Contiguous Block)
+---------+        +-------------------------------+
| len (N) |        | [Obj 0] [Obj 1] ... [Obj N] | 
| cap (M) |        +-------------------------------+
| ptr   --+------>  0x1400...
+---------+

### The unsafe Magic

The core logic relies on converting pointers to integers (uintptr) to perform arithmetic calculations that Go forbids by default.

Example of manual address calculation:
```Go

// Calculating the address of index 'i'
// Address = BasePtr + (Index * SizeOf(T))
targetAddr := unsafe.Pointer(uintptr(v.ptr) + (uintptr(i) * elementSize))

Preventing Memory Leaks (Loitering Objects)

In languages with Garbage Collection (like Go), simply decrementing the length in PopBack is dangerous if the vector holds pointers. The referenced objects would remain in memory because the underlying array still holds the address.

This implementation solves it by explicitly zeroing out the memory:
Go

func (v *Vector[T]) PopBack() {
    v.len--
    // ... calculate address ...
    
    // Zero out the memory to release GC reference
    var zero T
    *(*T)(targetAddr) = zero 
}
```
### 🛠 API Reference

    NewVect[T](cap): Creates a new vector with initial capacity.

    PushBack(val): Adds an element to the end. Reallocates if necessary.

    PopBack(): Removes the last element and clears memory to avoid leaks.

    Size(): Returns current number of elements.

    Capacity(): Returns current reserved memory size.