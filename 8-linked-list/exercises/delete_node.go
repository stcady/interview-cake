package main

import "fmt"

func deleteNode(nodeToDelete *LinkedListNode) error {
	nextNode := nodeToDelete.next
	if nextNode != nil {
		nodeToDelete.value = nextNode.value
		nodeToDelete.next = nextNode.next
	} else {
		return fmt.Errorf("attempting to delete last node")
	}
	return nil
}

type LinkedListNode struct {
	value interface{}
	next  *LinkedListNode
}
