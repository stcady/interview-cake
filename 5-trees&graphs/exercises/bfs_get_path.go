package main

import (
	"fmt"
)

func bfsGetPath(graph map[string][]string, startNode string, endNode string) ([]string, error) {

	// check if nodes are in the graph
	if _, ok := graph[startNode]; !ok {
		return []string{}, fmt.Errorf("startNode not in graph")
	}
	if _, ok := graph[endNode]; !ok {
		return []string{}, fmt.Errorf("endNode not in graph")
	}

	// create and initialzie queue
	nodesToVisit := []string{}
	nodesToVisit = append(nodesToVisit, startNode)

	// create and initialize visitation map
	howWeReachedNode := make(map[string]string)
	howWeReachedNode[startNode] = ""

	// iterate through each node to visit in the queue
	for len(nodesToVisit) > 0 {

		// pop left from queue
		currentNode := nodesToVisit[0]
		nodesToVisit = nodesToVisit[1:]

		if currentNode == endNode {
			// reconstruct the path
			return reconstructPath(howWeReachedNode, startNode, endNode), nil
		}

		// iterate through each connecting node from the current node
		for _, neighbor := range graph[currentNode] {

			// check if the node has already been visited
			if _, ok := howWeReachedNode[neighbor]; !ok {
				// add node to queue if it has not been visited
				nodesToVisit = append(nodesToVisit, neighbor)
				// add node to visited nodes
				howWeReachedNode[neighbor] = currentNode
			}
		}
	}

	// no path to end node
	return []string{}, nil
}

func reconstructPath(paths map[string]string, startNode string, endNode string) []string {
	shortestPath := []string{}
	currentNode := endNode
	// iterate through the paths in reverse
	for currentNode != "" {
		shortestPath = append(shortestPath, currentNode)
		currentNode = paths[currentNode]
	}
	return reverse(shortestPath)
}

func reverse(list []string) []string {
	i := 0
	j := len(list) - 1
	for i < j {
		list[i], list[j] = list[j], list[i]
		i += 1
		j -= 1
	}
	return list
}
