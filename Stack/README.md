# Go Stack (Thread-Safe LIFO)

A robust, thread-safe, and flexible implementation of a Stack (Last-In, First-Out) data structure in Go.

Designed for concurrent environments, this data structure uses `sync.RWMutex` to ensure safe access across multiple goroutines while maintaining high performance. It abstracts the underlying implementation via an interface, making it ideal for managing state, "undo" operations, and concurrent algorithms.

## 🚀 Features

- **Thread-Safe**: Built-in `sync.RWMutex` protection ensures safe concurrent access by multiple producers and consumers without panic.
- **Interface-Based**: Exposed via a clean `Stack` interface, allowing you to decouple implementation details from your application logic.
- **Memory Safe**: Automatically handles slice allocation and boundary checks to prevent runtime errors.
- **Dynamic**: Supports storing `any` type (mixed data) or can be easily adapted to Generics for strict typing.

## 📦 Installation

You can simply copy the `stack.go` file into your project, or import it if hosted as a module:

```bash
go get [github.com/JustJ3di/Golletions/Stack](https://github.com/JustJ3di/Golletions/Stack)

## 📖 Usage

```go
package main

import (
    "fmt"
    "log"

    stack "[github.com/JustJ3di/Golletions/Stack](https://github.com/JustJ3di/Golletions/Stack)"
)

func main() {
    // 1. Initialize (returns the Stack interface)
    s := stack.NewStack[string]()

    // 2. Push items (LIFO order)
    s.Push("First")
    s.Push("Second")
    s.Push(42) // Supports mixed types via 'any'

    // 3. Pop items
    val, err := s.Pop()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(val) // Output: 42 (The last item inserted)

    val, _ = s.Pop()
    fmt.Println(val) // Output: "Second"
    
    // 4. Handle Empty Stack
    _, err = s.Pop()
    _, err = s.Pop()
    val, err = s.Pop()
    if err != nil {
        fmt.Println("Stack is empty!") // Output: Stack is empty!
    }
}

```