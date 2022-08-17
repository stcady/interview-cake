package main

func Reverse(head *LinkedListNode) *LinkedListNode {
	currentNode := head
	previousNode := &LinkedListNode{}
	nextNode := &LinkedListNode{}
	for currentNode != (&LinkedListNode{}) {
		nextNode = currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	return previousNode
}

type LinkedListNode struct {
	value interface{}
	next  *LinkedListNode
}
