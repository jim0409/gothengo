package main

/*
Given a Binary Search Tree and a target number,
return true if there exist two elements in the BST
such that their sum is equal to the given target.

Example 1:
Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9
Output: True


Example 2:
Input:
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28
Output: False
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {

	set := make(map[int]int)
	return find(root, k, set)

}

func find(root *TreeNode, k int, set map[int]int) bool {

	if root == nil {
		return false
	}
	if _, ok := set[k-root.Val]; ok {
		return ok
	}

	set[root.Val]++
	return find(root.Left, k, set) || find(root.Right, k, set)

}

func main() {

}

/*
solution:
- https://leetcode.com/problems/two-sum-iv-input-is-a-bst/discuss/373456/Go-golang-clean-solution

func findTarget(root *TreeNode, k int) bool {

    set := make(map[int]int)
    return find(root, k, set)

}

func find(root *TreeNode, k int, set map[int]int) bool {

    if root == nil {
        return false
    }
    if _, ok := set[k - root.Val]; ok {
        return ok
    }

    set[root.Val]++
    return find(root.Left, k, set) || find(root.Right, k, set)

}
*/
