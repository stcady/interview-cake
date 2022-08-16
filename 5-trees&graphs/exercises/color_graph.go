package main

func colorGraph(graph Graph, colors []int) bool {

	for _, node := range graph.nodes {

		// check that graph is not cyclic
		if check(node.neighbors, node.id) {
			return false
		}

		// set illegal colors
		illegalColors := []int{}
		for _, neighbor := range node.neighbors {
			illegalColors = append(illegalColors, graph.nodes[neighbor].color)
		}

		// check colors
		for _, color := range colors {
			if !check(illegalColors, color) {
				node.color = color
				break
			}
		}
	}
	return true
}

func check(list []int, item int) bool {
	for _, element := range list {
		if element == item {
			return true
		}
	}
	return false
}

type Graph struct {
	nodes []*GraphNode
}

type GraphNode struct {
	id        int
	neighbors []int
	color     int
}

func (g *Graph) AddNode() (id int, color int) {
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
