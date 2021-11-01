package main

/*
Given the root of a binary tree, return the preorder traversal of its nodes' values.

Example 1:
1
 \
  2
 /
3
Input: root = [1,null,2,3]
Output: [1,2,3]

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
Output: [1,2]

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

// https://leetcode.com/problems/binary-tree-preorder-traversal/discuss/422814/Go-recursive-solution
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	preOrder := make([]int, 0)
	preOrder = append(preOrder, root.Val)
	preOrder = append(preOrder, preorderTraversal(root.Left)...)
	preOrder = append(preOrder, preorderTraversal(root.Right)...)

	return preOrder
}

func main() {}

/*
refer:
- https://leetcode.com/problems/binary-tree-preorder-traversal/discuss/748582/Go-golang-clean-solutions


Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
Memory Usage: 2 MB, less than 88.57% of Go online submissions for Binary Tree Preorder Traversal.

func preorderTraversal(root *TreeNode) []int {
    ans := []int{}
    helper(root, &ans)
    return ans
}

func helper(root *TreeNode, ans *[]int) {
    if root == nil { return }
    *ans = append(*ans, root.Val)
    helper(root.Left, ans)
    helper(root.Right, ans)
}



Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Binary Tree Preorder Traversal.

func preorderTraversal(root *TreeNode) []int {
    stack, ans := []*TreeNode{}, []int{}
    for root != nil || len(stack) != 0 {
        if root == nil {
            root = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
        }
        ans = append(ans, root.Val)
        if root.Right != nil {
            stack = append(stack, root.Right)
        }
        root = root.Left
    }
    return ans
}
*/
