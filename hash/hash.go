package main

import "fmt"

const ArraySize int = 5

type HashTable struct {
	Array [ArraySize]*HeadNode
}

type HeadNode struct {
	Head *Node
}

func createHeadNode() *HeadNode {
	return &HeadNode{}
}

type Node struct {
	Key string
	Value string
	Next *Node
}

func createNode(k, v string) *Node {
	return &Node{
		Key: k,
		Value: v,
	}
}

func (n *Node) changeValue (v string) {
	n.Value = v
}

func (n *Node) searchKey(k string) (string, bool) {
	if n.Key == k {
		return n.Value, true
	}
	return "", false
}

func (hN *HeadNode) checkingDuplicate (k, v string) bool {
	if hN.Head == nil {
		return false
	}
	currentNode := hN.Head 
	for currentNode != nil {
		if currentNode.Key == k {
			currentNode.changeValue(v)
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

func (hN *HeadNode) addHead(n *Node) {
	if hN.Head == nil {
		hN.Head = n
		return
	}

	n.Next = hN.Head
	hN.Head = n

}

func (hN *HeadNode) delete(k string) bool {
	if hN.Head == nil {
		return false
	}

	currentNode := hN.Head

	if currentNode.Key == k {
			hN.Head = hN.Head.Next
			return true
	}

	for currentNode.Next != nil {
		if currentNode.Next.Key == k {
			currentNode.Next = currentNode.Next.Next
			return true
		}
		currentNode = currentNode.Next
	}
	return false

}

func (hN *HeadNode) SearchKey(k string) (string, bool) {
	currentNode := hN.Head;
	for currentNode != nil {
		v , isFound := currentNode.searchKey(k)
		if isFound {
			return v, isFound
		}
		currentNode = currentNode.Next
	}
	return "", false
}


func CreateHashTable () *HashTable {
	hash := &HashTable{}
	for i, _ := range hash.Array {
		hash.Array[i] = createHeadNode()
	}
	return hash
}

func (h *HashTable) hash (totalAscii int) int {
	return  totalAscii % ArraySize
}

func (h *HashTable) sumAscii (k string) int {
	sum := 0 
	for _, v := range k {
		sum += int(v)
	}
	return sum
}

func (h *HashTable) Search (k string) (string,bool) {
	hashIndex := h.hash(h.sumAscii(k))
	v ,isExist := h.Array[hashIndex].SearchKey(k)
	if isExist {
		return  v, isExist
	}
	return v, isExist
}

func (h *HashTable) Insert(k, v string) {
	hashIndex := h.hash(h.sumAscii(k));
	if h.Array[hashIndex].checkingDuplicate(k,v) {
		return
	}
	newNode := createNode(k, v);
	h.Array[hashIndex].addHead(newNode)
}

func (h *HashTable) Delete (k string) bool {
	hashIndex := h.hash(h.sumAscii(k));
	return h.Array[hashIndex].delete(k)
}

func main () {

	hash := CreateHashTable();
	// fmt.Println(hash)
	hash.Insert("Syihab", "OMAGAAAAA")
	fmt.Println(hash.Array[3].Head)
	index, isExist := hash.Search("Apo")
	fmt.Println(index,isExist)


}
