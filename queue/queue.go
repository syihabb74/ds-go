package main

import "fmt"

type Queue struct {
	Items []int
}

func (q *Queue) Enqueue (item int) {
	fmt.Printf("%p\n", q.Items)
	q.Items = append(q.Items, item)
}

func (q *Queue) Dequeue () (int, bool) {
	fmt.Printf("%p\n", q.Items)
	if len(q.Items) == 0 {
		return 0, false
	}
	dequeue_item := q.Items[0];
	q.Items = q.Items[1:]
	return dequeue_item, true
}

func main () {

	queue := new(Queue);
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	fmt.Println(queue)
	// queue.Dequeue()
	// queue.Dequeue()
	// queue.Dequeue()
	// queue.Dequeue()
	// queue.Dequeue()
	// queue.Dequeue()
	// queue.Dequeue()
	// fmt.Println(queue)



}