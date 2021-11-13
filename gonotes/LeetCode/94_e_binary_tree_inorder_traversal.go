package main

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:
1
 \
  2
 /
3
Input: root = [1,null,2,3]
Output: [1,3,2]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]

Example 4:
  1
 /
2
Input: root = [1,2]
Output: [2,1]

Example 5:
1
 \
  2
Input: root = [1,null,2]
Output: [1,2]


Constraints:

The number of the nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

Follow up: Recursive solution is trivial, could you do it iteratively?
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	inOrder := make([]int, 0)
	inOrder = append(inOrder, inorderTraversal(root.Left)...)
	inOrder = append(inOrder, root.Val)
	inOrder = append(inOrder, inorderTraversal(root.Right)...)

	return inOrder
}

func main() {}

/*
refer:
- https://leetcode.com/problems/binary-tree-inorder-traversal/discuss/428608/Go-golang-0ms-solutions

Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Binary Tree Inorder Traversal.

func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    helper(root, &res)
    return res
}

func helper(root *TreeNode, s *[]int) {
    if root == nil { return }
    helper(root.Left, s)
    *s = append(*s, root.Val)
    helper(root.Right, s)
}


Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Binary Tree Inorder Traversal.

func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    cur := root
    stack := []*TreeNode{}
    for cur != nil || len(stack) > 0 {
        for cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
        }
        cur = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, cur.Val)
        cur = cur.Right
    }
    return res
}
*/
