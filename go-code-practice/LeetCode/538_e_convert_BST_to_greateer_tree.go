package main

/*
Greater Tree: 累加樹
Given a Binary Search Tree (BST), convert it to a Greater Tree such that
every key of the original BST is changed to the original key plus sum of all keys
greater than the original key in BST.

Example:

Input: The root of a Binary Search Tree like this:
              5
            /   \
           2     13

Output: The root of a Greater Tree like this:
             18
            /   \
          20     13
Note:
This question is the same as 1038:
- https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	helper(root, 0)
	return root
}

func helper(node *TreeNode, sum int) int {
	if node == nil {
		return sum
	}
	// 節點值為 原節點值 + helper(右節點值總和)
	node.Val += helper(node.Right, sum)
	return helper(node.Left, node.Val)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/convert-bst-to-greater-tree/discuss/100534/Simple-and-Clean-Go-Solution-Using-postorder-traversal


var sum int
func convertBST(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    sum = 0
    traverse(root)
    return root
}

func traverse(root *TreeNode) {
    if root == nil {
        return
    }
    traverse(root.Right)
    root.Val += sum
    sum = root.Val
    traverse(root.Left)
}
*/

/*
solution:
- https://leetcode.com/problems/convert-bst-to-greater-tree/discuss/642153/go-reverse-inorder

func convertBST(root *TreeNode) *TreeNode {
    sum := 0
    visit(root, &sum)
    return root
}

func visit(root *TreeNode, sum *int) {
    if root == nil {
        return
    }

    visit(root.Right, sum)
    root.Val += *sum
    *sum = root.Val
    visit(root.Left, sum)
}
*/
