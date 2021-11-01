package avl_tree

import (
	"fmt"
	"math/rand"
	"testing"
)

// Displays nodes left-depth first (used for debugging)
func (n *AVLNode) displayNodesInOrder() {
	if n.left != nil {
		n.left.displayNodesInOrder()
	}
	fmt.Print(n.key, " ")
	if n.right != nil {
		n.right.displayNodesInOrder()
	}
}

func (n *AVLNode) search(key int) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.key {
		return n.left.search(key)
	} else if key > n.key {
		return n.right.search(key)
	} else {
		return n
	}
}

type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Add(key int, value int) {
	t.root = t.root.add(key, value)
}

func (t *AVLTree) Remove(key int) {
	t.root = t.root.remove(key)
}

func (t *AVLTree) Update(oldKey int, newKey int, newValue int) {
	t.root = t.root.remove(oldKey)
	t.root = t.root.add(newKey, newValue)
}

func (t *AVLTree) Search(key int) *AVLNode {
	return t.root.search(key)
}

func (t *AVLTree) DisplayInorder() {
	t.root.displayNodesInOrder()
}

const (
	opAdd    = iota // 0
	opRemove        // 1
	opSearch        // 2
)

func TestAvlTree(t *testing.T) {
	tree := &AVLTree{}
	m := make(map[int]int)

	const maxKey = 100
	const nops = 10000

	for i := 0; i < nops; i++ {
		op := rand.Intn(3) // 0, 1, 2
		k := rand.Intn(maxKey)

		// consider an random walk condition
		switch op {
		case opAdd:
			v := rand.Int()
			tree.Add(k, v)
			m[k] = v
		case opRemove:
			tree.Remove(k)
			delete(m, k)
		case opSearch:
			var tv int
			node := tree.Search(k)
			tok := node != nil
			if tok {
				tv = node.Value
			}

			mv := m[k]
			if tv != mv {
				t.Errorf("Incorrect value for key %d, want: %d, got: %d", k, mv, tv)
			}
		}
	}
}
