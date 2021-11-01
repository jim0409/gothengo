package main

import "fmt"

/*

Given an array where elements are sorted in ascending order,

convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree

in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5],

which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5

*/

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	median := len(nums) / 2
	return &TreeNode{
		Val:   nums[median],
		Left:  sortedArrayToBST(nums[:median]),
		Right: sortedArrayToBST(nums[median+1:]),
	}
}

func main() {
	intArr := []int{3, 9, 2}
	fmt.Println(sortedArrayToBST(intArr))
}

/*
solution:
- https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/discuss/344138/Go-Recursive-solution

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	median := len(nums) / 2
	return &TreeNode{
		Val:   nums[median],
		Left:  sortedArrayToBST(nums[:median]),
		Right: sortedArrayToBST(nums[median+1:]),
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
