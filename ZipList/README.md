# Ziplist

ziplist is a memory-efficient, binary-encoded list implementation for Go. It stores various data types sequentially within a single byte slice ([]byte), reducing memory overhead compared to standard Go slices of interfaces.
Features

    Memory Efficient: Stores data in a compact binary format using Little Endian encoding.

    Type Support: Supports Integers (uint8 to uint64, int8 to int64), Floats, Booleans, and Strings.

    CRUD Operations: Support for Push, At (Get), Remove, and Clear.

## Usage

```go

package main

import (
	"fmt"
	"your-module/ziplist"
)

func main() {
	// Initialize with initial capacity
	zl := ziplist.New(128)

	// Add items
	zl.Push(uint32(100))
	zl.Push("Hello World")
	zl.Push(true)

	// Retrieve item at index 1
	val, err := zl.At(1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value: %v\n", val) // Output: Hello World

	// Remove item at index 0
	zl.Remove(0)
}
```
## Internal Structure

The ziplist uses a header to track total bytes and item count, followed by sequentially encoded entries:

[Total Bytes (4)] [Item Count (4)] [Entry 1] [Entry 2] ... [End Marker]

Each entry contains a Type byte, optional Length, and the Data.