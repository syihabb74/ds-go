package main

import (
	"errors"
	"fmt"
)

type Graph struct {
	Nodes []*Node
}

type Node struct {
	Value int
	Edges []*Node
}

func createNode(v int) *Node {
	return &Node{Value: v}
}

func createGraph() *Graph {
	return &Graph{}
}

func (g *Graph) AddNode(v int) {
	if len(g.Nodes) == 0 {
		g.Nodes = append(g.Nodes, createNode(v))
		return
	}

	if _ , isContain := g.ContainsNode(v); isContain {
		err := errors.New("existing value");
		fmt.Println(err.Error())
		return
	}

	g.Nodes = append(g.Nodes, createNode(v))

}

func (g *Graph) ContainsNode(v int) (int, bool) {
	for i, n := range g.Nodes {
		if n.Value == v {
			return i, true
		}
	}
	return -1, false
}

func (n *Node) ContainsEdge(v int) (int ,bool) {
	for i, n := range n.Edges {
		if n.Value == v {
			return i , true
		}
	}
	return -1, false
} 

func (g *Graph) ShowNodeWithEdge() {
	for _, n := range g.Nodes {
		fmt.Printf("%v -> [", n.Value)
		for j , e := range n.Edges {
			if j > 0 {
				fmt.Print(", ")
			} 
			fmt.Printf("%v", e.Value)
		}
		fmt.Println("]")
	}
}

func (g *Graph) AddEdge(from , to int) {
	iFrom , isContainFrom := g.ContainsNode(from);
	iTo, isContainTo := g.ContainsNode(to);
	if !isContainFrom || !isContainTo {
		fmt.Println("One of the node is not exist")
		return
	}

	if _, isEdgeExist := g.Nodes[iFrom].ContainsEdge(to) ; isEdgeExist {
		fmt.Println("Edge already exist")
		return
	}

	g.Nodes[iFrom].Edges = append(g.Nodes[iFrom].Edges, g.Nodes[iTo])

}






func main () {

	g := createGraph();
	g.AddNode(1);
	g.AddNode(2);
	g.AddNode(3);
	// g.AddNode(2);
	// g.ShowNodeWithEdge()
	// fmt.Println(g);
	g.AddEdge(1,2)
	g.AddEdge(2,1)
	g.AddEdge(1,3)
	g.ShowNodeWithEdge()

}