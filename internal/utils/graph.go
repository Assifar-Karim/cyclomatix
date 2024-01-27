package utils

import (
	"fmt"
	"strings"

	"github.com/dominikbraun/graph"
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

func (g Graph) CountEdges() int {
	result := 0
	for _, list := range g.AdjList {
		result += len(list)
	}
	return result
}

func (g Graph) GenerateDot() graph.Graph[string, string] {
	cteBlocksCounter := []int{0, 0, 0, 0, 0}
	dg := graph.New[string, string](graph.StringHash, graph.Directed())
	for node := range g.AdjList {
		n := sanitizeNodeName(strings.Replace(g.NodesNames[node], "\"", "\\\"", -1), cteBlocksCounter)
		g.NodesNames[node] = n
		dg.AddVertex(n)
	}

	for node, adjList := range g.AdjList {
		n := g.NodesNames[node]
		for _, adjNode := range adjList {
			aN := g.NodesNames[adjNode]
			dg.AddEdge(n, aN)
		}
	}
	return dg
}

func sanitizeNodeName(nodeName string, cteBlocksCounter []int) string {
	if strings.Contains(nodeName, "If End") {
		cteBlocksCounter[0]++
		return fmt.Sprintf("%v %v", nodeName, cteBlocksCounter[0])
	} else if strings.Contains(nodeName, "For End") {
		cteBlocksCounter[1]++
		return fmt.Sprintf("%v %v", nodeName, cteBlocksCounter[1])
	} else if strings.Contains(nodeName, "Switch End") {
		cteBlocksCounter[2]++
		return fmt.Sprintf("%v %v", nodeName, cteBlocksCounter[2])
	} else if strings.Contains(nodeName, "continue") {
		cteBlocksCounter[3]++
		return fmt.Sprintf("%v %v", nodeName, cteBlocksCounter[3])
	} else if strings.Contains(nodeName, "break") {
		cteBlocksCounter[4]++
		return fmt.Sprintf("%v %v", nodeName, cteBlocksCounter[4])
	} else {
		return nodeName
	}
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
