# IMPLEMENTATION SINGLE-LL & 

## Main COMPONENT
- Head (Representing starting point a linked list (where is node start))
- Tail (Representing End of a linked list) => Tail also can be representing by a nil node in the in -> &Node{Next : nil}
- Value (Representing a value it can be anything what you want to store just store it in this field)

## EXPLANATION
Single linked list is a structure data of linked list it's a group of nodes that are connected to each other and only in one direction, unable to return, only able to go to the next node and the next node is representing by pointer (memory address) -> like a chain.


### BENEFITS
Linked list data structure is very flexible in a size it can be small or big depend on what we define it on implementation

## TEST CASE OF WHEN SHOULD USE LINKED LIST
- PUSHING ELEMENT INTO FIRST SEQUENCE OF DATA (Because when you use this to slice it will be recomputate cause the sequence of element changing)
- WHEN THE DATA WILL BE VERY ACTIVE TO INSERT OR DELETING DATA 


```go

    type Node struct {
        Next *Node // Referencing Next Node via Memory Address
        Value int // Value (Representing a value it can be anything what you want to store just store it in this field)
        
       
    }

    type LinkedList struct {
        Head *Node // representing starting point of your Group node
        Tail *Node // ending point of your linked list 
         // if you wan't tail to mark ending point of a linked list and also
        // if you add this approach to your linked list it can be accelerate pushing element into the end
        // instead using slice because when you use slice with full cap and add new element
        // it will be creating a new memory in your computer and adding again one by one of existing element before you add new one
        // and it's not efficient
        // Tail also can be representing by a nil node in the in -> &Node{Next : nil}
        // OPTIONAL VALUE
        Length int // it's optional if you want to get how many element in your linked list without looping approach you can 
        // add this fiel
    }

 ```


