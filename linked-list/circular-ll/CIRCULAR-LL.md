# CIRCULAR LINKED-LIST


## LINKED-LIST COMPONENT

- Head (Representing starting point a linked list (where is node start))
- Tail (Representing End of a linked list) => Tail also can be representing by a nil node in the in -> &Node{Next : Head} it's because circular linked list is like ring there is no end point when it reaches the ending point it will be back to head e.g starting point
- Value (Representing a value it can be anything what you want to store just store it in this field)
- Length -> OPTIONAL to check how many element in your linked list instead of check the length by loop you can do this to accelerate counting element in your linked list


- Value You can put anything in this field as long as the value is data that you want to put in this field 
- Next -> Next Node --> this will be representing by memory address
- Prev -> Previous Node --> this will be representing by memory address

## ADDITIONAL BENEFITS FROM SINGLE LINKED LIST & DOUBLY LINKED LIST
- Make Accessing from the tail to the head is possible and faster without reverse traversing from the tail 

## CODE IMPLEMENTATION IN GO-LANG
``` GO
    type LinkedList struct {
        Head *Node
        Tail *Node // -> Tail.Next = LinkedList.Head
        Length int // OPTIONAL FIELD
    }

    type Node struct {
        Next *Node
        Prev *Node
        Value any
    }

    /* 
        actually the implementation of the method is freely on your hand 
        if you want add prefend and append method in your LinkedList struct 
        it's no problem or you want add the real impelemntation on Node struct itself it can be and the LinkedList method just facilitate the changes of the Head / Tail With Simple method
    */
    

```