package main

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head *Node
	Tail *Node
	Length int
}

type Node struct {
	Next *Node
	Prev *Node
	Value int
}



func NewNode (val int) *Node {
	return &Node{
		Next: nil,
		Prev: nil,
		Value: val,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
		Length: 0,
	}
}

func checkNodeInput (newNode *Node) error {

	if newNode == nil {
		return errors.New("invalid input")
	} 
	// fmt.Println("Gak error dari check node input")
	return nil

}

func (n *Node) Prefend (newNode *Node) (*Node, error) {

	err := checkNodeInput(newNode)
	if err != nil {
		return nil, err
	}
	// fmt.Println(newNode)
	n.Prev = newNode;
	// fmt.Println("Gak error dari prefend Node")
	n.Prev.Next = n
	return newNode, nil

}

func (n *Node ) Append (newNode *Node) (*Node, error) {

	err := checkNodeInput(newNode);
	if err != nil {
		return nil, err
	}

	n.Next = newNode
	n.Next.Prev = n
	return  newNode , nil

}

func (dll *LinkedList) Initialize (newNode *Node) {

		dll.Head = newNode;
		dll.Tail = newNode
		dll.Length++;

}

func (dll *LinkedList) Prefend (newNode *Node) {
	
	if dll.Head == nil {
		dll.Initialize(newNode)
		return;
	}

	newNode, err := dll.Head.Prefend(newNode)

	if err != nil {
		fmt.Println(err.Error())
		return
	}


	dll.Head = newNode
	dll.Length++;

}

func (dll *LinkedList ) Append (newNode *Node) {

	if dll.Head == nil {
		dll.Initialize(newNode)
		return;
	}

	newNode, err := dll.Tail.Append(newNode)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dll.Tail = newNode
	dll.Length++

}

func (dll *LinkedList) RemoveHead () {

	if dll.Head == nil {
		fmt.Println("Linked list is empty");
		return
	}

	dll.Head = dll.Head.Next
	if dll.Head == nil {
		dll.Tail = nil
		dll.Length--
		return
	}
	dll.Head.Prev = nil
	dll.Length--

}

func (dll *LinkedList) RemoveTail () {

	if dll.Head == nil {
		fmt.Println("Linked list is empty");
		return
	}

	dll.Tail = dll.Tail.Prev
	if dll.Tail == nil {
		dll.Head = nil
		dll.Length--
		return
	}
	dll.Tail.Next = nil
	dll.Length--

}

func CatchError () {
	err := recover()
	if err != nil {

		fmt.Println(err)
	}
}



func main () {
	// defer CatchError()
	
	linkedList := NewLinkedList();
	node1 := NewNode(1);
	node2 := NewNode(2);
	node3 := NewNode(3);
	linkedList.Prefend(node1);
	linkedList.Prefend(node2);
	linkedList.Append(node3);
	linkedList.RemoveHead()
	// fmt.Println(linkedList)
	// fmt.Printf("pointer node2 = %p\n", node2)
	// fmt.Printf("pointer node2 from Tail  = %p\n", linkedList.Tail.Prev);
	// fmt.Printf("pointer node1 from Head  = %p\n", linkedList.Head.Next);
	fmt.Printf("%+v\n", linkedList.Head);
	// fmt.Printf("%+v\n", linkedList.Head.Next);
	fmt.Printf("%+v\n", linkedList.Tail);
	

}