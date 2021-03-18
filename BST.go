package main

import "fmt"

// Node is for the binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the binary search tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		//right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}

}

// Search will take in a key value
// and RETURN true if there is a node with that value
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(200)
	fmt.Println(tree)
        
	fmt.Println(tree.Search(50))

}