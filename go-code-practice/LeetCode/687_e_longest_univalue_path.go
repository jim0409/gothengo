package main

/*
Given a binary tree,
find the length of the longest path where each node in the path has the same value.
This path may or may not pass through the root.

The length of path between two nodes is represented by the number of edges between them.

Example 1:
Input:

              5
             / \
            4   5
           / \   \
          1   1   5
Output: 2


Example 2:
Input:

              1
             / \
            4   5
           / \   \
          4   4   5
Output: 2


Note:
The given binary tree has not more than 10000 nodes.
The height of the tree is not more than 1000.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/longest-univalue-path/discuss/491532/Go-98-solution
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	_ = helper(root, &max)
	return max
}

func helper(node *TreeNode, max *int) int {
	v, summary := 0, 0
	if node.Left != nil {
		vl := helper(node.Left, max)
		if node.Val == node.Left.Val {
			v = vl + 1
			summary += vl + 1
		}
	}
	if node.Right != nil {
		vr := helper(node.Right, max)
		if node.Val == node.Right.Val {
			if vr+1 > v {
				v = vr + 1
			}
			summary += vr + 1
		}
	}
	if summary > *max {
		*max = summary
	}
	return v
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/longest-univalue-path/discuss/493408/go-clean-code-64ms-beats-100


func longestUnivaluePath(root *TreeNode) int {
	var ret int
	helper(root, 0, &ret)
	return ret
}

func helper(node *TreeNode, prev int, ret *int) int {
	if node == nil {
		return 0
	}
	left, right := helper(node.Left, node.Val, ret), helper(node.Right, node.Val, ret)
	*ret = int(math.Max(float64(*ret), float64(left+right)))
	if node.Val == prev {
		return 1 + int(math.Max(float64(left), float64(right)))
	}
	return 0
}
*/
