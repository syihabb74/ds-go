# HEAP 

## HEAP COMPONENT 
- SLICE  -> of any data types can be compare


## BENEFITS
- Fast retrieval of prioritized data
- Efficient insertion and removal
- Ideal for priority-based processing

```go 
    type Heap struct {
	Slice []int
    } 

    func (h *Heap) Heapify (v int) {
	// insert your implemtation here
}


## METHOD AND FUNCTION

func (h *Heap) MaxHeapifyUp (i int) { 
    // insert your implemtation here
}

func (h *Heap) MaxHeapifyDown (i int) {
    // insert your implemtation here
}

func (h *Heap) swap (i1, i2 int) {
	// insert your implemtation here
}

func (h *Heap) left (i int) int {
	// insert your implemtation here
}

func (h *Heap) right (i int) int {
	// insert your implemtation here
}


func (h *Heap) Extract () int {
    // insert your implemtation here
}
    func parent(index int) int { 
        // insert your implemtation here
    }

    ```