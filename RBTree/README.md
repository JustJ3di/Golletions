## Red-Black Tree (`RBtree`)

The `RBtree` allows you to store values associated with a key.  
The key must satisfy the `Key` interface (ordered, comparable types such as `int`, `float`, or `string`).

```go
package main

import (
	"fmt"

	rbtree "github.com/JustJ3di/Golletions/RBTree"
)
func main() {
    tree := rbtree.New[int]()

    tree.InsertKeyValue(10, "Apple")
    tree.InsertKeyValue(5,  "Banana")
    tree.InsertKeyValue(20, "Cherry")

    val := tree.Search(10)
    fmt.Printf("Key 10: %v\n", val)

    fmt.Printf("Min: %v\n", tree.Min())
    fmt.Printf("Max: %v\n", tree.Max())

    deleted := tree.Delete(5)
    if deleted {
        fmt.Println("Key 5 deleted successfully.")
    }

    fmt.Println("Current Tree State:")
    tree.PrintInOrder()
}
```

---

## 📚 API Reference

### `RBtree[T]`

| Method | Description | Complexity |
|------|------------|------------|
| `New[T]()` | Creates a new empty Red-Black Tree | `O(1)` |
| `InsertKeyValue(k T, v any)` | Inserts a new key-value pair | `O(log n)` |
| `Delete(k T)` | Removes the node with the specified key | `O(log n)` |
| `Search(k T)` | Returns the value associated with key `k` | `O(log n)` |
| `Min()` | Returns the value of the minimum key | `O(log n)` |
| `Max()` | Returns the value of the maximum key | `O(log n)` |
| `Clear()` | Removes all nodes from the tree | `O(1)` |

---

## 🗺 Roadmap

- [x] Red-Black Tree  
- [ ] AVL Tree  
- [ ] Graph (Adjacency List / Matrix)  
- [ ] Trie (Prefix Tree)  
- [ ] BTree and B+Tree

---