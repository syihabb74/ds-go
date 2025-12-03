package main

import "fmt"

type Root struct {
	RootNode *Node
	Height int
}

type Node struct {
	Left *Node
	Right *Node
	Val int
}

func (n *Node) LeftChild (v int) {
	if n.Left != nil && n.Right != nil {
		n.Left.LeftChild(v)
	} else if n.Left == nil {
		n.Left = newNode(v)
	} else {
		n.Right = newNode(v)
	}
}

func (n *Node) RightChild (v int) {
	if n.Right != nil && n.Left != nil {
		n.Right.RightChild(v)
	} else if n.Right == nil {
		n.Right = newNode(v)
	} else {
		n.Left = newNode(v)
	}
}

func newNode (v int) *Node {
	return &Node{Val: v}
}

func rootCreate (value int) *Root {

	return &Root{RootNode: newNode(value)}

}

func (r *Root) addLeft (v int) {
	if r.RootNode.Left == nil {
		r.RootNode.Left = newNode(v)
		return
	} else {
		r.RootNode.Left.LeftChild(v)
	}
}

func (r *Root) addRight (v int) {
	if r.RootNode.Right == nil {
		r.RootNode.Right = newNode(v)
		return
	} else {
		r.RootNode.Right.RightChild(v)
	}
}

func main () {

	root := rootCreate(3);
	root.addLeft(5);
	root.addLeft(7);
	root.addLeft(8);
	root.addRight(11)
	fmt.Printf("%+v\n",root.RootNode)

}