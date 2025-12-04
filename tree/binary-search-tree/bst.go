package main

import "fmt"

type Root struct {
	RootNode *Node
}

func CreateRootNode(value int) *Root {

	return &Root{RootNode: newNode(value)}

}

func (r *Root) InsertChild(value int) {
	r.RootNode.Insert(value)
}

func (r *Root) SearchValue(value int) bool {

	return r.RootNode.Search(value)

}

type Node struct {
	Value int // you can define it anything for example string
	Left  *Node
	Right *Node
}

func newNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func (n *Node) Insert(value int) {

	if value >= n.Value {
		if n.Right == nil {
			n.Right = newNode(value)
		} else {
			n.Right.Insert(value)
		}
	} else {
		if n.Left == nil {
			n.Left = newNode(value)
		} else {
			n.Left.Insert(value)
		}
	}

}

func (n *Node) Search(value int) bool {

	if value == n.Value {
		return true
	} else if value > n.Value {
		if n.Right == nil {
			return false
		}
		return n.Right.Search(value)
	} else {
		if n.Left == nil {
			return false
		}
		return n.Left.Search(value)
	}

}

func main() {
	root := CreateRootNode(5);
	root.RootNode.Insert(9)
	root.RootNode.Insert(1)
	root.RootNode.Insert(2)
	root.RootNode.Insert(3)
	root.RootNode.Insert(6)
	root.RootNode.Insert(7)
	root.RootNode.Insert(8)
	root.RootNode.Insert(10)
	fmt.Println(root.RootNode.Right)
	fmt.Println(root.SearchValue(100))
}
