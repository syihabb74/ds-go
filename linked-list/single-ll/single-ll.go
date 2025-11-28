package main

import (
	"fmt"
	// "testing"
)






type Node struct {
	Next *Node
	Value int
}

type LinkedList struct {
	Head *Node
	Length int /* 
	this is Optional if you want to sacrifices to counting total length
	in single linked-list but it sacrifices a litle memory alloc to your 
	memory
	*/

}

func (hN *LinkedList) Prefend (n *Node) {
	if hN.Head == nil {
		hN.Head = n
		hN.Length++	
		return;
	}
	curHeadNode := hN.Head
	// fmt.Println(curHeadNode)
	hN.Head = n
	// fmt.Printf("%+v\n", *hN.Head)
	hN.Head.Next = curHeadNode
	// fmt.Printf("%+v\n", *hN.Head)
	hN.Length++
}

func (hN *LinkedList) PrintVal () {
	if hN.Length == 0 {
		return;
	}
	curNode := hN.Head;
	fmt.Println(curNode, "CurNOde")
	// fmt.Println(curNode.Node)
	for {
		fmt.Println(*curNode)
		curNode = curNode.Next
		if curNode == nil {
			break
		}
	}
}

// implementing linked list with Head Tail and dyn Val
type Node2 struct {
	Next *Node2
	Value any // it can be anything
}

func (n2 *Node2) Push (newNode *Node2) {

	n2.Next = newNode

}

type LinkedList2 struct {
	Head *Node2
	Tail *Node2
	Length int
}

func NewLinkedList2 () LinkedList2 {

	return LinkedList2{};

}

func (ll2 *LinkedList2) Prefend (newNode *Node2) {

	if ll2.Head == nil {
		ll2.Head = newNode
		ll2.Tail = newNode
		ll2.Length++
		return;
	}

	curHead := ll2.Head
	ll2.Head = newNode
	ll2.Head.Push(curHead)
	ll2.Length++
	return;

}


func (ll2 *LinkedList2) Prafend (newNode *Node2) {

	if ll2.Head == nil {
		fmt.Println("Can't add tail before add Node to the head")
		return;
	}
	ll2.Tail.Next = newNode
	ll2.Tail = newNode
	ll2.Length++
	return

}


func (ll2 *LinkedList2) PrintVENode2 () {
	curNode := ll2.Head;
	for curNode != nil {
		fmt.Println(curNode.Value);
		curNode = curNode.Next
	}
}

func (ll2 *LinkedList2) RemoveHeadLl2 () {
	if ll2.Head == nil {
		fmt.Println("It's empty Linked List Can't remove ")
		return;
	}

	if ll2.Tail == nil {
		ll2.Head = nil
		ll2.Length--
		return;
	}
	// fmt.Println("Masuk sini???")
	curNode := ll2.Head;
	if curNode.Next == nil {
		ll2.Head = nil;
		ll2.Tail = nil
	} else {
		ll2.Head = curNode.Next
	}
	// fmt.Println(curNode.Next)
	ll2.Length--
	return



}

func NewNode2 (v any) *Node2 {

	return &Node2{Value: v}

}



func main () {

	// single_ll := LinkedList{}

	// // fmt.Printf("%+v\n", single_ll);
	// n1 := &Node{Value: 1};
	// n2 := &Node{Value: 2};
	// single_ll.Prefend(n1);
	// single_ll.Prefend(n2)
	// // fmt.Printf("%+v\n", *single_ll.Head);
	// single_ll.PrintVal()

	// ---------


	


	node1 := NewNode2(1);
	node2 := NewNode2(2);
	// fmt.Println(node1);
	ll2 := NewLinkedList2();
	// fmt.Println(ll2)
	ll2.Prefend(node1)
	ll2.Prafend(node2)
	// fmt.Println(ll2)
	// fmt.Println(node1);
	// fmt.Println(ll2.Tail)
	// fmt.Println(&node2)

	// test  := "aaaa"

	// var str *string = &test
	// fmt.Println("string =====> ",str)

	// testNum := 1212;

	// var num *int = &testNum
	// fmt.Println("num ====> ",num)


	// var slice []*int = []*int{
	// 	num,
	// }

	// fmt.Println(slice)

	ll2.PrintVENode2()
	ll2.RemoveHeadLl2()
	ll2.PrintVENode2()
	fmt.Printf("%+v\n",ll2)

}