package rbtree

import (
	"fmt"
)

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

func (n *rbnode[T]) grandparent() *rbnode[T] {
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

func (rb *RBtree[T]) Clear() {
	rb.root = nil
}

func (rb *RBtree[T]) search(k T) *rbnode[T] {
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
func (rb *RBtree[T]) Search(k T) any {
	if n := rb.search(k); n != nil {
		return n.value
	}
	return nil
}

func (rb *RBtree[T]) Insert(k T, v any) {

	n := &rbnode[T]{
		parent: nil,
		left:   nil,
		right:  nil,
		col:    red, //Default
		key:    k,
		value:  v,
	}
	rb.insert(n)

}

func (rb *RBtree[T]) insert(n *rbnode[T]) {
	if try := rb.search(n.key); try != nil {
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
		if n.parent == n.grandparent().left {
			u := n.grandparent().right //uncle
			if u != nil && u.col == red {
				n.parent.col = black
				u.col = black
				n.grandparent().col = red
				n = n.grandparent()
			} else {
				if n == n.parent.right {
					n = n.parent
					rb.leftRotate(n)
				}
				n.parent.col = black
				n.grandparent().col = red
				rb.rightRotate(n.grandparent())
			}
		} else {
			u := n.grandparent().left //uncle
			if u != nil && u.col == red {
				n.parent.col = black
				u.col = black
				n.grandparent().col = red
				n = n.grandparent()
			} else {
				if n == n.parent.left {
					n = n.parent
					rb.rightRotate(n)
				}
				n.parent.col = black
				n.grandparent().col = red
				rb.leftRotate(n.grandparent())
			}
		}

	}
	rb.root.col = black
}

func (rb *RBtree[T]) Min() any {
	curr := rb.root
	for curr.left != nil {
		curr = curr.left
	}
	return curr.value
}

func (rb *RBtree[T]) Max() any {
	curr := rb.root
	for curr.right != nil {
		curr = curr.right
	}
	return curr.value
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

func (rb *RBtree[T]) transplant(u, v *rbnode[T]) {
	if u.parent == nil {
		rb.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

/*
Delete a RBnode.
Follow the clrs algorithm
*/
func (rb *RBtree[T]) Delete(k T) bool {

	min := func(n *rbnode[T]) *rbnode[T] {
		for n.left != nil {
			n = n.left
		}
		return n
	}

	z := rb.search(k)
	if z == nil {
		return false // No node found
	}

	var x *rbnode[T]
	y := z
	yOriginalColor := y.col

	// We need to track x's parent explicitly because if x is nil,
	// we cannot access x.parent in the fixup function.
	var xParent *rbnode[T]

	if z.left == nil {
		x = z.right
		xParent = z.parent // x replaces z, so its parent becomes z's parent
		rb.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		xParent = z.parent
		rb.transplant(z, z.left)
	} else {
		y = min(z.right)
		yOriginalColor = y.col
		x = y.right

		if y.parent == z {
			xParent = y // x is the child of y
		} else {
			xParent = y.parent // x was y's child
			rb.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}

		rb.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.col = z.col
	}

	if yOriginalColor == black {
		rb.deleteFixup(x, xParent)
	}

	return true
}

func (rb *RBtree[T]) deleteFixup(n, parent *rbnode[T]) {

	for n != rb.root && (n == nil || n.col == black) {
		if n == parent.left {
			w := parent.right // Sibling

			// Case 1: Sibling is Red
			if w.col == red {
				w.col = black
				parent.col = red
				rb.leftRotate(parent)
				w = parent.right
			}

			// Case 2: Sibling's children are both Black (or nil)
			if (w.left == nil || w.left.col == black) && (w.right == nil || w.right.col == black) {
				w.col = red
				n = parent
				parent = n.parent
			} else {
				// Case 3: Sibling's right child is Black (Left is Red)
				if w.right == nil || w.right.col == black {
					if w.left != nil {
						w.left.col = black
					}
					w.col = red
					rb.rightRotate(w)
					w = parent.right
				}

				// Case 4: Sibling's right child is Red
				w.col = parent.col
				parent.col = black
				if w.right != nil {
					w.right.col = black
				}
				rb.leftRotate(parent)
				n = rb.root // Terminate
			}
		} else {
			// Mirror images of the cases above (n is a right child)
			w := parent.left

			// Case 1 (Mirror)
			if w.col == red {
				w.col = black
				parent.col = red
				rb.rightRotate(parent)
				w = parent.left
			}

			// Case 2 (Mirror)
			if (w.right == nil || w.right.col == black) && (w.left == nil || w.left.col == black) {
				w.col = red
				n = parent
				parent = n.parent
			} else {
				// Case 3 (Mirror)
				if w.left == nil || w.left.col == black {
					if w.right != nil {
						w.right.col = black
					}
					w.col = red
					rb.leftRotate(w)
					w = parent.left
				}

				// Case 4 (Mirror)
				w.col = parent.col
				parent.col = black
				if w.left != nil {
					w.left.col = black
				}
				rb.rightRotate(parent)
				n = rb.root // Terminate
			}
		}
	}

	if n != nil {
		n.col = black
	}
}
