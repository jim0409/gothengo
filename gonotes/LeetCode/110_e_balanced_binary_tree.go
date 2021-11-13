package main

/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the left and right subtrees of every node differ in height by no more than 1.


Example 1:

Given the following tree [3,9,20,null,null,15,7]:

    3
   / \
  9  20
    /  \
   15   7
Return true.

Example 2:

Given the following tree [1,2,2,3,3,null,null,4,4]:

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
*/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lh := height(root.Left)
	rh := height(root.Right)

	if absDiff(lh, rh) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	count := 1 + max(height(node.Left), height(node.Right))

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absDiff(lh, rh int) int {
	res := lh - rh
	if res < 0 {
		return -res
	}
	return res
}

func main() {
	a := []int{1, 2, 3, 4, 5}
}

/*
solution:
https://leetcode.com/problems/balanced-binary-tree/discuss/380609/go-8ms

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lh := height(root.Left)
	rh := height(root.Right)

	if absDiff(lh, rh) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func absDiff(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxHeight(height(root.Left), height(root.Right))
}

func maxHeight(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
*/

/*
solution 2
https://leetcode.com/problems/balanced-binary-tree/discuss/344128/Go-Recursive-solution

func isBalanced(root *TreeNode) bool {
	return maxDepth(root) != -1
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
