package main

import "fmt"

/*
Given a binary tree,

check whether it is a mirror of itself
(ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return ish(root.Left, root.Right)
}

func ish(l, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == r
	}

	if l.Val != r.Val {
		return false
	}

	// 左邊樹的左節點 與 右邊樹的右節點 做比較
	// 左邊樹的右節點 與 右邊樹的左節點 做比較
	return ish(l.Left, r.Right) && ish(l.Right, r.Left)
}

func main() {
	intArr1 := []int{1, 2, 3, 4, 5}
	tree1 := ImplementTree(intArr1)
	fmt.Println(isSymmetric(tree1))
}

/*
solution:
- https://leetcode.com/problems/symmetric-tree/discuss/198893/Go-Recursive-4ms


func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return ish(root.Left, root.Right)
}

func ish(l, r *TreeNode) bool {
    if l == nil || r == nil {
        return l == r
    }

    if l.Val != r.Val {
        return false
    }

    return ish(l.Left, r.Right) && ish(l.Right, r.Left)
}
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ImplementTree(arr []int) *TreeNode {
	rtNode := TreeNode{arr[0], nil, nil}

	for i := 1; i < len(arr); i++ {
		rtNode.InsertNode(arr[i])
	}

	return &rtNode
}

func (n *TreeNode) InsertNode(v int) {

	if n == nil {
		n = &TreeNode{v, nil, nil}

	} else {

		if v < n.Val {
			if n.Left == nil {
				n.Left = &TreeNode{v, nil, nil}
			} else {
				n.Left.InsertNode(v)
			}
		} else {
			if n.Right == nil {
				n.Right = &TreeNode{v, nil, nil}
			} else {
				n.Right.InsertNode(v)
			}
		}
	}
}
