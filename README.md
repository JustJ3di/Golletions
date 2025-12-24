# Golletions 📦

**Go + Collections = Golletions**

Golletions is a library of efficient, generic data structures written in pure Go.  
It is designed to be **zero-dependency**, **type-safe** (using Go 1.18+ generics), and **easy to integrate** into any Go project.

Currently, the library provides a robust implementation of a **Red-Black Tree**, with more data structures planned for future releases.

---

## 🚀 Installation

```bash
go get github.com/yourusername/golletions
```

---

## ✨ Features

- **Pure Go**  
  No Cgo or external dependencies.

- **Generics**  
  Fully type-safe implementations using Go 1.18+ generics.

- **Red-Black Tree**  
  A self-balancing binary search tree supporting  
  **O(log n)** insert, delete, and search operations.

- **Key Support**  
  Supports standard ordered types (`int`, `uint`, `float`, `string`) via a custom constraints interface.

---

## 🛠 Usage

### Red-Black Tree (`RBtree`)

The `RBtree` allows you to store values associated with a key.  
The key must satisfy the `Key` interface (ordered, comparable types such as `int`, `float`, or `string`).

```go
package main

import (
    "fmt"
    "github.com/yourusername/golletions/rbtree"
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

## 🤝 Contributing

Contributions are welcome!

1. Fork the project  
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)  
3. Commit your changes (`git commit -m "Add some AmazingFeature"`)  
4. Push to the branch (`git push origin feature/AmazingFeature`)  
5. Open a Pull Request  

---

## 📄 License

Distributed under the **MIT License**.  
See the `LICENSE` file for more information.
