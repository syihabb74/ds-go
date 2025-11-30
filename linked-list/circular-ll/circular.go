package main
import "fmt"

type CircularLL struct {
	Head *Node
	Tail *Node
	// Length int OPTIONAL VALUE
}


type Node struct {
	Next *Node
	Prev *Node
	Value any // if you want to spesify what data type you will use just 
	// write any keyword with explicit data type you want int, string, or another struct
	// it's freely 
}

func NewNode (value any) *Node {
	return &Node{Value: value}
}


func NewCircularLL () *CircularLL {

	return &CircularLL{}

}



func (CLL *CircularLL) Prefend (newNode *Node) {

	if newNode == nil {
		fmt.Printf("invalid input \n")
		return;
	}

	if CLL.Head == nil {
		CLL.Head = newNode
		CLL.Tail = newNode
		newNode.Prev = newNode
		newNode.Next = newNode
		return
	}


	if CLL.isSingleNode() {
		newNode.Next = CLL.Tail
		newNode.Prev = CLL.Tail
		CLL.Tail.Next = newNode
		CLL.Tail.Prev = newNode
		CLL.Head = newNode
		return
	}

	// fmt.Println(CLL.isSingleNode())

	 newNode.Next = CLL.Head
	 newNode.Prev = CLL.Tail
	 CLL.Head.Prev = newNode
	 CLL.Tail.Next = newNode
	 CLL.Head = newNode


}

func (CLL *CircularLL ) Append (newNode *Node) {

	if newNode == nil {
		fmt.Printf("invalid input \n")
		return;
	}

	fmt.Println(CLL)
	if CLL.Head == nil {
		fmt.Println("Masuk KE appen")
		// fmt.Println("Masuk sekali")
		CLL.Head = newNode
		CLL.Tail = newNode
		newNode.Prev = newNode
		newNode.Next = newNode
		// fmt.Println(*CLL)
		return
	}

	if CLL.isSingleNode() {
		// fmt.Println("Masuk sekali ke single node")
		CLL.Head.Next = newNode
		CLL.Head.Prev = newNode
		newNode.Next = CLL.Head
		newNode.Prev = CLL.Head
		CLL.Tail = newNode
		// fmt.Println(CLL)
		return
	}

	// fmt.Println("Masuk sini")

	newNode.Prev = CLL.Tail
	newNode.Next = CLL.Head
	// fmt.Println("New node reassign",newNode.Next)
	CLL.Tail.Next = newNode
	CLL.Head.Prev = newNode
	CLL.Tail = newNode
	// fmt.Println("INI HARUsnya SATU",CLL.Tail.Next)

}

func (CLL *CircularLL) RemoveHead () {

	if CLL == nil {
		fmt.Println("Element is empty");
		return;
	}

	if CLL.Head.Next == CLL.Head.Prev {
		CLL.Head = nil
		CLL.Tail = nil
		return
	}

	CLL.Head = CLL.Head.Next
	CLL.Head.Prev = CLL.Tail
	CLL.Tail.Next = CLL.Head

}

func (CLL *CircularLL) RemoveTail () {
	if CLL == nil {
		fmt.Println("Element is empty");
		return;
	}
	if CLL.Tail.Next == CLL.Tail.Prev {
		CLL.Head = nil
		CLL.Tail = nil
		return
	}
	CLL.Tail = CLL.Tail.Prev
	CLL.Tail.Next = CLL.Head
	CLL.Head.Prev = CLL.Tail
}


func (CLL *CircularLL) isSingleNode () bool {

	return CLL.Head == CLL.Tail

}



func checkingType (a *Node, b *Node) bool {
	return a == b
}


func main () {

	cll := NewCircularLL();
	n1 := NewNode(1);
	n2 := NewNode(2);
	// n3 := NewNode(3);
	n4 := NewNode(0)
	cll.Prefend(n1);
	cll.Append(n2);
	// cll.Prefend(n3);
	cll.Append(n4)
	cll.RemoveTail()
	// ptr1 := fmt.Sprintf("%p", cll.Head)
	// ptr2 := fmt.Sprintf("%p", cll.Tail);
	// fmt.Printf("%v\n", ptr1)
	// fmt.Printf("%v\n", ptr2)
	// checkingPtr := checkingType(cll.Head, cll.Tail)
	// fmt.Printf("%+v\n", cll)

	fmt.Printf("%+v\n", cll)
	fmt.Printf("HEAD\n%+v\n", cll.Head)
	// fmt.Printf("Next Node Dari Cll HEAD\n%+v\n", cll.Head.Next)
	// fmt.Printf("Next Node dari Cll HEAD.Next\n%+v\n", cll.Head.Next.Next.Next)
	// fmt.Printf("Next NOde dari Cll TAIL\n%+v\n", cll.Tail)
	fmt.Printf("TAIL\n%+v\n", cll.Tail)

}


// HEAD
// &{Next:0x140000700a0 Prev:0x140000700a0 Value:1}
// Next Node Dari Cll HEAD
// &{Next:0x140000700c0 Prev:0x14000070080 Value:2}
// Next NOde dari Cll TAIL
// &{Next:0x14000070080 Prev:0x140000700a0 Value:0}


// 0x140000700a0 = 2
// 0x140000700a0 = 2

// 0x140000700c0
// 0x140000700a0


// &{Next:0x14000122080 Prev:0x140001220a0 Value:1}
// Next Node Dari Cll HEAD
// &{Next:0x140001220a0 Prev:0x14000122060 Value:2}
// Next NOde dari Cll TAIL
// &{Next:0x14000122060 Prev:0x14000122080 Value:0}

// 0x14000122080 = 2
// 0x14000122080 = 2

// 0x140001220a0 = 0
// 0x140001220a0 = 0

// 0x14000122060 = 1
// 0x14000122060 = 1








