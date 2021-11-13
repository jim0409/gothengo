package main

import "fmt"

/*
Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical

and the nodes have the same value.


Example 1:

Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true



Example 2:

Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false



Example 3:

Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false

*/

// use recursion to solve
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	intArr1 := []int{1, 2, 3, 4, 5}
	intArr2 := []int{1, 2, 3, 4}
	tree1 := ImplementTree(intArr1)
	tree2 := ImplementTree(intArr2)
	fmt.Println(isSameTree(tree1, tree2))

}

func createBinaryTree(arr []int) *TreeNode {
	return nil
}

/*
solution:
- https://leetcode.com/problems/same-tree/discuss/299782/Go-Solution


func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
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
