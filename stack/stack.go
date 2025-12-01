package main

import "fmt"

type Stack struct {
	Items []int
}

func (s *Stack) Push (i int) {
	s.Items = append(s.Items, i)
}	

func (s *Stack) Pop () (int, bool) {

	l := len(s.Items);

	if l == 0 {
		return 0, false
	}

	item_popped := s.Items[l-1];
	s.Items = s.Items[:l-1];
	return item_popped, true


	
}



func main () {

	stack := Stack{Items: []int{1,2,3}};
	stack.Pop();
	stack.Pop();
	stack.Pop();
	stack.Pop();
	fmt.Println(stack);
	

}