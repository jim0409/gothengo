package main

/*
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its minimum depth = 2.
*/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 || right == 0 {
		return 1 + max(left, right)
	}
	return 1 + min(left, right)
}
func min(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}
func max(x, y int) int {
	if x <= y {
		return y
	} else {
		return x
	}
}

func main() {

}

/*
solution :
https://leetcode.com/problems/minimum-depth-of-binary-tree/discuss/164985/Go-solution

func minDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }
    left := minDepth(root.Left)
    right := minDepth(root.Right)
    if left == 0 || right == 0 {
        return 1 + max(left, right)
    }
    return 1 + min(left, right)
}
func min(x, y int) int{
    if x >= y{
        return y
    }else{
        return x
    }
}
func max(x, y int) int{
    if x <= y{
        return y
    }else{
        return x
    }
}
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
