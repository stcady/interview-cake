package main

func isBalanced(tree *BinaryTree) bool {
	root := tree.root
	if root == nil {
		return true
	}
	depths := []int{}
	nodes := []Pair{}
	nodes = append(nodes, Pair{root, 0})
	nodesLength := 1
	for nodesLength > 1 {
		node, depth := nodes[len(nodes)-1].a, nodes[len(nodes)-1].b
		nodes = nodes[:len(nodes)-1]
		if node.left == nil && node.right == nil {
			if !check(depths, depth) {
				depths = append(depths, depth)
				if (len(depths) > 2) || (len(depths) == 2 && Abs(depths[0]-depths[1]) > 1) {
					return false
				}
			}
		} else {
			if node.left != nil {
				nodes = append(nodes, Pair{node.left, depth + 1})
			}
			if node.right != nil {
				nodes = append(nodes, Pair{node.right, depth + 1})
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

type Pair struct {
	a *BinaryNode
	b int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
