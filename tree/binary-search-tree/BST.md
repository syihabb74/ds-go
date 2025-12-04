# BINARY SEARCH TREE (BST)

## NODE COMPONENT
- VALUE -> Value can be anything string, number everything can be compare each other is valid
- LEFT -> Location element is more smaller than the parent 
- RIGHT -> Location element is more bigger than the parent



## RULES
- When the value we insert and the value that already exist in the node same there is another rule
    - it can be left child alias more smaller than parent node
    - it can be right child alias more bigger than parent node
    - but usually it will go to the right side

## BENEFIT
- To make more efficient and accelerate your searching algorithm O(log n) time complexity

## NOTE
- If the implementation of BST applied to array/slice make sure the array/slice is sorted 


    ```go 
    type Root struct {
	RootNode *Node
    } // this the Root nodes if you don't want to make it you can go straight to Node struct 

    type Node struct {
	Value int // you can define it anything for example string
	Left  *Node
	Right *Node
    } 

    func (n *Node) Insert(value int) {
        // insert your implementation here
    }

    func (n *Node) Search(value int) {
        // insert your implementation here
    }

    ```