package main

/*
A binary tree is univalued if every node in the tree has the same value.

Return true if and only if the given tree is univalued.


Example 1:
     1
   /   \
  1     1
 / \   / \
1   1 1   1

Input: [1,1,1,1,1,null,1]
Output: true


Example 2:
     2
   /   \
  2     2
 / \
5   2

Input: [2,2,2,5,2]
Output: false


Note:
The number of nodes in the given tree will be in the range [1, 100].
Each node's value will be an integer in the range [0, 99].
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/univalued-binary-tree/discuss/346408/Go-golang-0ms-clean-solution
func isUnivalTree(root *TreeNode) bool {

	var a, b bool

	if root.Left == nil || root.Left.Val == root.Val && isUnivalTree(root.Left) {
		a = true
	}
	if root.Right == nil || root.Right.Val == root.Val && isUnivalTree(root.Right) {
		b = true
	}
	return a && b
}

func main() {}

/*
refer:
- https://leetcode.com/problems/univalued-binary-tree/discuss/536680/Go-golang-simple-0ms-solution

func DFS(root *TreeNode, v int) bool {
    if root == nil {
        return true
    }
    return root.Val == v && DFS(root.Left, root.Val) && DFS(root.Right, root.Val)
}
func isUnivalTree(root *TreeNode) bool {
    return DFS(root, root.Val)
}
*/
