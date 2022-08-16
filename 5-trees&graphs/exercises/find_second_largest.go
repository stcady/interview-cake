package main

func findSecondLargest(root *BinaryNode) int {
	if root == nil || (root.left == nil && root.right == nil) {
		return root.data
	}
	current := root
	for current != nil {
		if current.left != nil && current.right == nil {
			return findLargest(current.left)
		}
		if current.right != nil && current.right.left == nil && current.right.right == nil {
			return current.data
		}
		current = current.right
	}
	return 0
}

func findLargest(root *BinaryNode) int {
	current := root
	for current != nil {
		if current.right == nil {
			return current.data
		}
		current = current.right
	}
	return 0
}

type BinaryNode struct {
	data  int
	left  *BinaryNode
	right *BinaryNode
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}
