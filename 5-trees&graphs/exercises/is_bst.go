package main

func isBst(root *BinaryNode, lowerBound int, upperBound int) bool {
	if root == nil {
		return true
	}
	if root.data >= upperBound || root.data <= lowerBound {
		return false
	}
	return isBst(root.left, lowerBound, root.data) && isBst(root.right, root.data, upperBound)
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
