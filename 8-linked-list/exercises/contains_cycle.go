package main

func containsCycle(firstNode *LinkedListNode) bool {
	slowIterator := firstNode
	fastIterator := firstNode
	for fastIterator != nil && fastIterator.next != nil {
		slowIterator = slowIterator.next
		fastIterator = fastIterator.next.next
		if fastIterator == slowIterator {
			return true
		}
	}
	return false
}

type LinkedListNode struct {
	value interface{}
	next  *LinkedListNode
}
