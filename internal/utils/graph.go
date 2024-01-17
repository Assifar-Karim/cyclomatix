package utils

import (
	"fmt"
)

type Graph struct {
	AdjList    map[int32][]int32
	NodesNames []string
	LatestNode int32
}

func NewCFGraph() *Graph {
	nodesNames := []string{"start"}
	adjList := make(map[int32][]int32)
	adjList[0] = []int32{}

	return &Graph{
		AdjList:    adjList,
		NodesNames: nodesNames,
		LatestNode: int32(len(adjList)) - 1,
	}
}

func (g *Graph) Append(node string) {
	g.NodesNames = append(g.NodesNames, node)
	latestNode := int32(len(g.AdjList)) - 1
	g.AdjList[latestNode], latestNode = append(g.AdjList[latestNode], latestNode+1), latestNode+1
	g.AdjList[latestNode] = []int32{}
	g.LatestNode = latestNode
}

func (g *Graph) AddStatement(stmt string) {
	g.NodesNames[g.LatestNode] += stmt + "\n"
}

// This method creates a new node without performing the appending on the adj list
func (g *Graph) AddNode(node string) {
	g.NodesNames = append(g.NodesNames, node)
	g.LatestNode++
	g.AdjList[g.LatestNode] = []int32{}
}

func (g *Graph) DeleteNode(node int32) {
	if node == g.LatestNode {
		g.LatestNode--
	}
	nodesNames := []string{}
	nodesNames = append(nodesNames, g.NodesNames[:node]...)
	g.NodesNames = append(nodesNames, g.NodesNames[node+1:]...)

	delete(g.AdjList, node)
	for k, v := range g.AdjList {
		for idx, n := range v {
			if n == node {
				replacement := []int32{}
				replacement = append(replacement, v[:idx]...)
				g.AdjList[k] = append(replacement, v[idx+1:]...)
				break
			}
		}
	}

}

func (g *Graph) LinkNodes(node1, node2 int32) {
	g.AdjList[node1] = append(g.AdjList[node1], node2)
}

// NOTE (KARIM) : This method is used for debugging purposes
func (g *Graph) Print() {
	fmt.Println(g.AdjList)
	for key, values := range g.AdjList {
		fmt.Printf("[%v] -> [%v]\n", key, values)
		for _, value := range values {
			fmt.Printf("[[%v]] ->[[%v]]\n", g.NodesNames[key], g.NodesNames[value])
		}

	}
}
