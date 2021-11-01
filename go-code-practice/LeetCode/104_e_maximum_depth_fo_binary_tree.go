package main

import "fmt"

/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path

from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.



Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.

*/

func maxDepth(root *TreeNode) int {
	// 假設節點為空，回傳0，否則conut++
	if root == nil {
		return 0
	}

	// 同理，root.Left & root.Right往下比較。並且回傳Max
	count := 1 + Max(maxDepth(root.Left), maxDepth(root.Right))
	return count

}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	intarr := []int{4, 9, 20, nil, nil, 15, 7}
	tree := ImplementTree(intarr)
	fmt.Println(maxDepth(tree))
}

/*
solution:
- https://leetcode.com/problems/maximum-depth-of-binary-tree/discuss/379551/Go-solution


func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }

    count := 1 + Max(maxDepth(root.Left),maxDepth(root.Right))
    return count

}

func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
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
