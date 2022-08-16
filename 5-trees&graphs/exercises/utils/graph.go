package main

type Graph struct {
	nodes []*GraphNode
}

type GraphNode struct {
	id        int
	neighbors []int
	color     string
}

func (g *Graph) AddNode() (id int, color string) {
	id = len(g.nodes)
	g.nodes = append(g.nodes, &GraphNode{
		id:        id,
		neighbors: []int{},
		color:     color,
	})
	return
}

func (g *Graph) Neighbors(id int) []int {
	neighbors := []int{}
	for _, node := range g.nodes {
		for neighbor := range node.neighbors {
			if node.id == id {
				neighbors = append(neighbors, neighbor)
			}
			if neighbor == id {
				neighbors = append(neighbors, node.id)
			}
		}
	}
	return neighbors
}

func (g *Graph) Nodes() []int {
	nodes := make([]int, len(g.nodes))
	for i := range g.nodes {
		nodes[i] = i
	}
	return nodes
}
