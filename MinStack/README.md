# Go MinStack (Thread-Safe LIFO)

A robust, thread-safe, and flexible implementation of a MinStack (Last-In, First-Out) data structure in Go.


## 📦 Installation


```bash
go get [github.com/JustJ3di/Golletions/Stack](https://github.com/JustJ3di/Golletions/MinStack)
```

## 📖 Usage

```go
package main

import (
    "fmt"

    st "[github.com/JustJ3di/Golletions/Stack](https://github.com/JustJ3di/Golletions/Stack)"
)

func main() {
	st := NewMinStack[int]()
	st.Push(12)
	st.Push(14)
	st.Push(20)
	st.Push(2)
	st.Push(100)

	fmt.Println(st.Min())
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
	a, _ := st.Pop()
	fmt.Println(st)
	fmt.Println("Popped = ", a)

	fmt.Println("New min = ", st.Min())

	fmt.Println(st)

}

```