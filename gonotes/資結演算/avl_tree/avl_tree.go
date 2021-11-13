package avl_tree

type AVLNode struct {
	key   int
	Value int

	// height counts nodes (not edges)
	height int
	left   *AVLNode
	right  *AVLNode
}

func (n *AVLNode) add(key int, value int) *AVLNode {
	if n == nil {
		return &AVLNode{key, value, 1, nil, nil}
	}

	if key < n.key {
		n.left = n.left.add(key, value)
	} else if key > n.key {
		n.right = n.right.add(key, value)
	} else {
		n.Value = value
	}

	return n.rebalanceTree()
}

func (n *AVLNode) remove(key int) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = n.left.remove(key)
	} else if key > n.key {
		n.right = n.right.remove(key)
	} else { // consider condition is delete key is right node key
		if n.left != nil && n.right != nil {
			rightMinNode := n.right.findSmallest()
			n.key = rightMinNode.key
			n.Value = rightMinNode.Value
			// delete the smallest key node from right node group
			n.right = n.right.remove(rightMinNode.key)
		} else if n.right == nil { // consdier n.right is nil
			n = n.left
		} else if n.left == nil {
			n = n.right
		} else {
			n = nil
			return n
		}
	}

	return n.rebalanceTree()
}

// search from the left node and find the most left nodes
func (n *AVLNode) findSmallest() *AVLNode {
	if n.left != nil {
		return n.left.findSmallest()
	}

	return n
}

// according to wiki: rebalance once height isn't within -1 0 1
func (n *AVLNode) rebalanceTree() *AVLNode {
	if n == nil {
		return n
	}
	// calculate the n node height
	n.recalculateHeight()

	balanceFactor := n.left.getHeight() - n.right.getHeight()

	if balanceFactor == -2 { // consider right skew more
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 { // consider left skew more
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}

	return n
}

/*
1(n)					2
  2(n')			==>   1   3
    3
*/
func (n *AVLNode) rotateLeft() *AVLNode {
	newRoot := n.right
	n.right = newRoot.left // consider `2` exist left node child
	newRoot.left = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

/*
	3(n)				2
  2(n')			==>   1   3
1
*/
func (n *AVLNode) rotateRight() *AVLNode {
	newRoot := n.left
	n.left = newRoot.right // consider `2` exist right node child
	newRoot.right = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

func (n *AVLNode) recalculateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

func (n *AVLNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
