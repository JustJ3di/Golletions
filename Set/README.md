# Go Set (Thread-Safe)

A robust, thread-safe, and flexible implementation of a Set  data structure in Go.

Designed for concurrent environments, this data structure uses `sync.RWMutex` to ensure safe access across multiple goroutines while maintaining high performance. It abstracts the underlying implementation via an interface, making it ideal for managing state, "undo" operations, and concurrent algorithms.

## 🚀 Features

- **Thread-Safe**: Built-in `sync.RWMutex` protection ensures safe concurrent access by multiple producers and consumers without panic.
- **Interface-Based**: Exposed via a clean `Set` interface, allowing you to decouple implementation details from your application logic.
- **Memory Safe**: Automatically handles slice allocation and boundary checks to prevent runtime errors.
- **Dynamic**: Supports storing `any` type (mixed data) or can be easily adapted to Generics for strict typing.

## 📦 Installation

You can simply copy the `Set.go` file into your project, or import it if hosted as a module:

```bash
go get [github.com/JustJ3di/Golletions/Stack](https://github.com/JustJ3di/Golletions/Set)
```

## 📖 Usage

```go
package main

import (
    "fmt"
    "log"

    stack "[github.com/JustJ3di/Golletions/Stack](https://github.com/JustJ3di/Golletions/Set)"
)


func main() {
	ms := NewSet[int]()

	ms.Add(1)
	ms.Add(3)
	fmt.Println(ms.Contains(2))

	ms.Clear()
}

```