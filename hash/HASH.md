# HASH DATA STRUCTURE

## HASH COMPONENT
- Array of *HeadNode
- Each index represents a bucket that stores a linked list of nodes (separate chaining)

## NODE COMPONENT
- Key   -> unique identifier for each value
- Value -> data stored in the hash table
- Next  -> pointer to the next node in case of collision (linked list)

## COLLISION HANDLING
- Uses separate chaining
- When collision occurs, new nodes are inserted at the head of the linked list

## BENEFITS
- Search complexity:
  - Average case: O(1)
  - Worst case: O(n) when many collisions occur in the same bucket



``` go 
    const ArraySize int = 5

type HashTable struct {
	Array [ArraySize]*HeadNode
}

type HeadNode struct {
	Head *Node
}

type Node struct {
	Key string
	Value string
	Next *Node
}

func createHeadNode() *HeadNode {
    // implement your code here
}

func createNode(k, v string) *Node {
    // implement your code here
}

func (n *Node) changeValue (v string) {
    // implement your code here
}

func (n *Node) searchKey(k string) (string, bool) {
    // implement your code here
}

func (hN *HeadNode) checkingDuplicate (k, v string) bool {
    // implement your code here
}

func (hN *HeadNode) addHead(n *Node) {
    // implement your code here
}

func (hN *HeadNode) delete(k string) bool {
    // implement your code here
}

func (hN *HeadNode) SearchKey(k string) (string, bool) {
    // implement your code here
}


func CreateHashTable () *HashTable {
    // implement your code here
}

func (h *HashTable) hash (totalAscii int) int {
	
}

func (h *HashTable) sumAscii (k string) int {
    // implement your code here
}

func (h *HashTable) Search (k string) (string,bool) {
    // implement your code here
}

func (h *HashTable) Insert(k, v string) {
    // implement your code here
}

func (h *HashTable) Delete (k string) bool {
    // implement your code here
}
```