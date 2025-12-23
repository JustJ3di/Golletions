package rbtree

import "fmt"

type Key interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type color bool

const (
	red   color = false
	black color = true
)

type rbnode[T Key] struct {
	key                 T
	value               any //don't touch, this field it must be modified only by the user
	col                 color
	parent, left, right *rbnode[T]
}

func (n *rbnode[T]) Grandparent() *rbnode[T] {
	if n != nil {
		return n.parent.parent
	}
	return nil
}

type RBtree[T Key] struct {
	root *rbnode[T]
}

func New[T Key]() *RBtree[T] {
	return &RBtree[T]{root: nil}
}

func (rb *RBtree[T]) Search(k T) *rbnode[T] {
	curr := rb.root
	for curr != nil && k != curr.key {
		if k > curr.key {
			curr = curr.right
		} else {
			curr = curr.left
		}
	}
	return curr
}
func (rb *RBtree[T]) SearchValue(k T) any {
	if n := rb.Search(k); n != nil {
		return n.value
	}
	return nil
}

func (rb *RBtree[T]) InsertKeyValue(k T, v any) {

	n := &rbnode[T]{
		parent: nil,
		left:   nil,
		right:  nil,
		col:    red, //Default
		key:    k,
		value:  v,
	}
	rb.Insert(n)

}

func (rb *RBtree[T]) Insert(n *rbnode[T]) {
	if try := rb.Search(n.key); try != nil {
		try.value = n.value
	} else if rb.root == nil {
		rb.root = n
		rb.root.col = black
	} else {
		curr := rb.root
		var parent *rbnode[T]
		for curr != nil {
			parent = curr
			if n.key > curr.key {
				curr = curr.right
			} else {
				curr = curr.left
			}
		}
		n.parent = parent
		if n.key < n.parent.key {
			n.parent.left = n
		} else {
			n.parent.right = n
		}

		n.col = red
		rb.fix(n)
	}
}

func (rb *RBtree[T]) rightRotate(n *rbnode[T]) {
	y := n.left

	n.left = y.right
	if y.right != nil {
		y.right.parent = n
	}

	y.parent = n.parent

	if n.parent == nil {
		rb.root = y
	} else if n == n.parent.right {
		n.parent.right = y
	} else {
		n.parent.left = y
	}

	y.right = n
	n.parent = y
}

func (rb *RBtree[T]) leftRotate(n *rbnode[T]) {
	y := n.right
	n.right = y.left

	if y.left != nil {
		y.left.parent = n
	}

	y.parent = n.parent

	if n.parent == nil {
		rb.root = y
	} else if n == n.parent.right {
		n.parent.right = y
	} else {
		n.parent.left = y
	}

	y.left = n
	n.parent = y

}

func (rb *RBtree[T]) fix(n *rbnode[T]) {
	for n != rb.root && n.parent.col == red {
		if n.parent == n.Grandparent().left {
			u := n.Grandparent().right //uncle
			if u != nil && u.col == red {
				n.parent.col = black
				u.col = black
				n.Grandparent().col = red
				n = n.Grandparent()
			} else {
				if n == n.parent.right {
					n = n.parent
					rb.leftRotate(n)
				}
				n.parent.col = black
				n.Grandparent().col = red
				rb.rightRotate(n.Grandparent())
			}
		} else {
			u := n.Grandparent().left //uncle
			if u != nil && u.col == red {
				n.parent.col = black
				u.col = black
				n.Grandparent().col = red
				n = n.Grandparent()
			} else {
				if n == n.parent.left {
					n = n.parent
					rb.rightRotate(n)
				}
				n.parent.col = black
				n.Grandparent().col = red
				rb.leftRotate(n.Grandparent())
			}
		}

	}
	rb.root.col = black
}

func (rb *RBtree[T]) PrintInOrder() {
	rb.orderedPrintRecursive(rb.root)
}

func (n *rbnode[T]) String() string {
	c := "Black"
	if n.col == red {
		c = "Red"
	}
	return fmt.Sprintf("{Key: %v, Val: %v, Color: %s}", n.key, n.value, c)
}

func (rb *RBtree[T]) orderedPrintRecursive(n *rbnode[T]) {
	if n != nil {
		rb.orderedPrintRecursive(n.left)
		fmt.Println(n)
		rb.orderedPrintRecursive(n.right)
	}

}

func (rb *RBtree[T]) delete(k T) bool {
	if rb.Search(k) == nil {
		return false //No node found
	}

	return true
}
