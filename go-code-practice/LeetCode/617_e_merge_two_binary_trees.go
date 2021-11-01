package main

/*
Given two binary trees and imagine that when you put one of them to cover the other,
some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree.

The merge rule is that
if two nodes overlap, then sum node values up as the new value of the merged node.
Otherwise, the NOT null node will be used as the node of new tree.

Example 1:
Input:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
Output:
Merged tree:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7


Note: The merging process must start from the root nodes of both trees.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/merge-two-binary-trees/discuss/104374/Simple-and-clean-7-lines-in-Go
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	return &TreeNode{t1.Val + t2.Val, mergeTrees(t1.Left, t2.Left), mergeTrees(t1.Right, t2.Right)}
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/merge-two-binary-trees/discuss/341451/Go-O(n)-with-explanation

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// Handle edge cases like:
	// t1: [], t2: [1]
	if t1 == nil {
		return t2
	}
	// t1: [1], t2: []
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val

	// If t1 left is empty then we append t2 left to t1.
	if t1.Left == nil && t2.Left != nil {
		t1.Left = t2.Left
	} else {
		// otherwise we just call the mergeTrees recursively.
		mergeTrees(t1.Left, t2.Left)
	}

	// If t1 right is empty then we append t2 right to t1.
	if t1.Right == nil && t2.Right != nil {
		t1.Right = t2.Right
	} else {
		// otherwise we just call the mergeTrees recursively.
		mergeTrees(t1.Right, t2.Right)
	}

	return t1
}
*/
