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

## 📚 API Reference

### `RBtree[T]`

| Method | Description | Complexity |
|------|------------|------------|
| `New[T]()` | Creates a new empty Red-Black Tree | `O(1)` |
| `Insert(k T, v any)` | Inserts a new key-value pair | `O(log n)` |
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
