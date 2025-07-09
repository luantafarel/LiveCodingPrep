package main

import "fmt"

type BinaryTreeFuncs interface {
	Insert()
	Search()
	Remove()
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func NewNode(key int) *Node {
	return &Node{Key: key}
}

func NewBST() *BST {
	return &BST{}
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if key == n.Key {
		return true
	} else if key < n.Key {
		return n.Left.Search(key)
	} else {
		return n.Right.Search(key)
	}
}

func (bst *BST) Search(key int) bool {
	return bst.Root.Search(key)
}

func (n *Node) Insert(key int) {
	if key < n.Key {
		if n.Left == nil {
			n.Left = NewNode(key)
		} else {
			n.Left.Insert(key)
		}
	} else {
		if n.Right == nil {
			n.Right = NewNode(key)
		} else {
			n.Right.Insert(key)
		}
	}
}

func (bst *BST) Insert(key int) {
	if bst.Root == nil {
		bst.Root = NewNode(key)
	} else {
		bst.Root.Insert(key)
	}
}

func (n *Node) Delete(key int) *Node {
	if n == nil {
		return nil
	}

	if key < n.Key {
		n.Left = n.Left.Delete(key)
	} else if key > n.Key {
		n.Right = n.Right.Delete(key)
	} else {
		if n.Left == nil && n.Right == nil {
			return nil
		}
		if n.Left == nil {
			return n.Right
		}
		if n.Right == nil {
			return n.Left
		}
		minRight := n.Right.minValueNode()
		n.Key = minRight.Key
		n.Right = n.Right.Delete(minRight.Key)
	}
	return n
}

func (n *Node) minValueNode() *Node {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (bst *BST) Delete(key int) {
	bst.Root = bst.Root.Delete(key)
}

func main() {
	bst := NewBST()

	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(8)

	fmt.Println("Search 7:", bst.Search(7)) // true
	fmt.Println("Search 9:", bst.Search(9)) // false

	bst.Delete(7)
	fmt.Println("Search 7 after delete:", bst.Search(7)) // false
}
