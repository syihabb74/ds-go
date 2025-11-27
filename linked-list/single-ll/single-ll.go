package main

import "fmt"

type Node struct {
	Node *Node
	Value int
}

type Head struct {
	Head *Node
	Length int /* 
	this is Optional if you want to sacrifices to counting total length
	in single linked-list but it sacrifices a litle memory alloc to your 
	memory
	*/

}

func (hN *Head) Prefend (n *Node) {
	if hN.Head == nil {
		hN.Head = n
		hN.Length++	
		return;
	}
	curHeadNode := hN.Head
	// fmt.Println(curHeadNode)
	hN.Head = n
	// fmt.Printf("%+v\n", *hN.Head)
	hN.Head.Node = curHeadNode
	// fmt.Printf("%+v\n", *hN.Head)
	hN.Length++
}

func (hN *Head) PrintVal () {
	if hN.Length == 0 {
		return;
	}
	curNode := hN.Head;
	fmt.Println(curNode, "CurNOde")
	// fmt.Println(curNode.Node)
	for {
		fmt.Println(*curNode)
		curNode = curNode.Node
		if curNode == nil {
			break
		}
	}
}


func main () {

	single_ll := Head{}

	// fmt.Printf("%+v\n", single_ll);
	n1 := &Node{Value: 1};
	n2 := &Node{Value: 2};
	single_ll.Prefend(n1);
	single_ll.Prefend(n2)
	// fmt.Printf("%+v\n", *single_ll.Head);
	single_ll.PrintVal()


}