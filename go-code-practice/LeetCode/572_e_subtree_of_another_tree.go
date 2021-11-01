package main

/*
Given two non-empty binary trees s and t,
check whether tree t has exactly the same structure
and node values with a subtree of s.

A subtree of s is a tree consists of a node in s
and all of this node's descendants.

The tree s could also be considered as a subtree of itself.

Example 1:
Given tree s:
     3
    / \
   4   5
  / \
 1   2

Given tree t:
   4
  / \
 1   2

Return true,
because t has the same structure and node values with a subtree of s.


Example 2:
Given tree s:
     3
    / \
   4   5
  / \
 1   2
    /
   0
Given tree t:
   4
  / \
 1   2
Return false.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isSameTree(s, t) {
		return true
	}
	if s == nil {
		return false
	}
	if isSubtree(s.Left, t) || isSubtree(s.Right, t) {
		return true
	}
	return false
}

func isSameTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return isSameTree(t1.Left, t2.Left) && isSameTree(t1.Right, t2.Right)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/subtree-of-another-tree/discuss/480180/Go-Tree-traversal-solution

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if isEqual(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isEqual(s *TreeNode, t *TreeNode) bool {
	if s == t {
		return true
	} else if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && isEqual(s.Left, t.Left) && isEqual(s.Right, t.Right)
}
*/
