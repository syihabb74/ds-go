package main

import (
	"fmt"
)

type GraphMatrix struct {
	NodeIndex map[rune]int
	Node []*Node
	Matrix [][]int
	size int
}

type Node struct {
	Value rune
}

func (g *GraphMatrix) AddNode(v rune) {
	if _, isExist := g.NodeIndex[v] ; isExist {
		fmt.Println("Node is existing")
		return;
	}
	node := createNode(v)
	g.Node = append(g.Node, node)
	g.NodeIndex[rune(v)] = len(g.Node) - 1
	g.size++
	g.Matrix = append(g.Matrix, make([]int, g.size))
	g.recreateMatrix(g.size)
}

func (g* GraphMatrix) recreateMatrix(size int) {
	for i := 0 ; i < size - 1 ; i++ {
		newSize := make([]int, size);
		copy(newSize, g.Matrix[i]);
		g.Matrix[i]= newSize
	}
}

func CreateMatrixGraph () *GraphMatrix {
	return &GraphMatrix{
		NodeIndex: make(map[rune]int),
	}
}

func (g *GraphMatrix) AddEdge(from, to rune) {
	iFrom, isExistFrom := g.NodeIndex[from];
	iTo, isExistTo := g.NodeIndex[to];
	if !isExistFrom || !isExistTo {
		fmt.Println("One of node is not exist");
		return;
	}

	if from == to {
		fmt.Println("Add self edge invalid")
		return
	}

	if g.Matrix[iFrom][iTo] == 1 {
		fmt.Println("Existing edge")
		return
	}

	g.Matrix[iFrom][iTo] = 1

}

func (g *GraphMatrix) PrintGraph () {
	for _, v1 := range g.Node {
		_ = v1
		fmt.Printf("%s",string(v1.Value))
		
		for _, v2 := range g.Node {
			// fmt.Printf("%s ",string(v2.Value))
			_ = v2
		}
		fmt.Println("")	
	}
	fmt.Println("")

}




func createNode(v rune) *Node {
	return &Node{
		Value: v,
	}
}

func main () {
	g1 := CreateMatrixGraph()
	g1.AddNode('A');
	g1.AddNode('B');
	g1.AddNode('C');
	g1.AddEdge('A', 'B')
	g1.AddEdge('A', 'C')
	g1.AddNode('D')
	// g1.AddEdge('A', 'B')
	// fmt.Printf("%+v\n", g1)
	g1.PrintGraph()



}