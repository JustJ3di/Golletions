# Go Trie (Prefix Tree)

A clean, efficient, and Unicode-compatible implementation of a Trie (Prefix Tree) data structure in Go.

This data structure is optimized for efficient string retrieval, making it ideal for autocomplete systems, spell checkers, and prefix-based matching.

## 🚀 Features

- **Efficiency**: Insert and Search operations are $O(L)$, where $L$ is the length of the string.
- **Unicode Support**: Uses `map[rune]*trienode` internally, allowing support for emojis, non-English scripts, and special characters.
- **Dynamic Growth**: Nodes are allocated only when needed using maps, saving memory on sparse datasets compared to fixed-array implementations.

## 📦 Installation

You can simply copy the `trie.go` file into your project, or if you are hosting this as a module.

---


```go
package main

import (
	"fmt"

	rbtree "github.com/JustJ3di/Golletions/Trie"
)

func main() {
    // 1. Initialize
    trie := gotrie.New()

    // 2. Insert words
    trie.Insert("apple")
    trie.Insert("app")
    trie.Insert("go")

    // 3. Search
    fmt.Println(trie.Search("apple")) // true
    fmt.Println(trie.Search("app"))   // true
    fmt.Println(trie.Search("ap"))    // false (exists as prefix, but not a whole word)
    fmt.Println(trie.Search("java"))  // false
}

```