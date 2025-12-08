package main

import (
	"fmt"
)

type Heap struct {
	Slice []int // you can put it with any type data as long the data can be compare
}

func (h *Heap) Heapify (v int) {
	h.Slice = append(h.Slice, v);
	h.MaxHeapifyUp(len(h.Slice) - 1)
	fmt.Println(h.Slice)
}

func (h *Heap) MaxHeapifyUp (i int) { // it must be input with the value len(h.Slice-1) at the beginning
	if len(h.Slice) < 1 {
		fmt.Println("Mac heapifying not working cause array is empty")
		return
	}
	childIndex := i // in the first we can say i as last index but following with the time 
					// and condition in for loop always true it will not break and child index
					// will be continue to change for example below : 
					/*
						if we have 4 element in the beginning we can say it's first actually the last index
						but the last index has more value than the parent it will swap and child index
						become the index 1 because we call parent method in the reassign expression
						childIndex = parent(childIndex) it's mean child will be change to previous parent
						and the value h.Slice(childIndex) become the child from the previous parent alias index 1
					*/

	for childIndex > 0 {
		if h.Slice[childIndex] > h.Slice[parent(childIndex)] {
			h.swap(parent(childIndex), childIndex)
			childIndex = parent(childIndex)
		} else {
			return;
		}
	}

}

func (h *Heap) MaxHeapifyDown (i int) {

	l, r := h.left(i), h.right(i) // this will be find the child for
								// left and right child of this index (parent)
								// at the beginning this will be [0]
	lastIndex := len(h.Slice) - 1;
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.Slice[l] > h.Slice[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		
		if h.Slice[childToCompare] > h.Slice[i] { // if the child more larger than parent below will be happening
			h.swap(i, childToCompare) // swap between parent and child to child become parent and parent become child
			i = childToCompare // i from 0 at the beginning will be follow which index has more value before swapping
			l, r = h.left(i), h.right(i) // l and r will be reassing to take child from the current element or new parent 
			// which old child index that has more value than old parent 
			/*
				for example : 
				0 -> formula curIndex * 2 + 1 / curIndex * 2 + 2 -> (1/2) will be parent 
				if the parent still smaller 
				1 -> formula curIndex * 2 + 1 / curIndex * 2 + 2 -> (3/4)  will be parent 
			*/
		} else {
			break
		}

	}



}

func (h *Heap) swap (i1, i2 int) {
	h.Slice[i1], h.Slice[i2] = h.Slice[i2], h.Slice[i1]
}

func (h *Heap) left (i int) int {
	return i * 2 + 1
}

func (h *Heap) right (i int) int {
	return i * 2 + 2
}

func parent(index int) int { // this function is use to search parent of current index
	return (index - 1) / 2
}

func (h *Heap) Extract () int {
	if len(h.Slice) == 0 {
		fmt.Println("Heap is empty")
		return -1
	}
	extracted := h.Slice[0];
	h.Slice[0] = h.Slice[len(h.Slice) - 1]
	h.Slice = h.Slice[:len(h.Slice) - 1]
	h.MaxHeapifyDown(0)
	return  extracted
}



func main () {

	
	
	heap := &Heap{};
	// heap.Heapify(10)
	// heap.Heapify(20)
	// heap.Heapify(30)
	// fmt.Println(heap)


	initialV := []int{
		30,1,2,13,15,9,11,5,10,7,
	}

	fmt.Println(initialV)

	for _, v := range initialV {
		heap.Heapify(v)
	}
	fmt.Println(heap)
	heap.Extract()
	heap.Extract()
	fmt.Println(heap)
	



}