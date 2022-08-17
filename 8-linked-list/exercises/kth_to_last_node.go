package main

import (
	"fmt"
)

func kthToLastNode(k int, head *LinkedListNode) (*LinkedListNode, error) {
	if k < 1 {
		return nil, fmt.Errorf("invalid length")
	}
	trailNode := head
	leadNode := head
	i := 0
	for i < k {
		leadNode = leadNode.next
		if leadNode == nil {
			return nil, fmt.Errorf("k is larger than list length")
		}
		i += 1
	}
	for leadNode.next != nil {
		trailNode = trailNode.next
		leadNode = leadNode.next
	}
	return trailNode, nil
}

type LinkedListNode struct {
	value interface{}
	next  *LinkedListNode
}
