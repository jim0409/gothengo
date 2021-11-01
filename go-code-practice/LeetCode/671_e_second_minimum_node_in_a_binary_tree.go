package main

/*
Given a non-empty special binary tree consisting of nodes with the non-negative value,
where each node in this tree has exactly two or zero sub-node.

If the node has two sub-nodes, then this node's value is the smaller value among its two sub-nodes.
More formally, the property root.val = min(root.left.val, root.right.val) always holds.

Given such a binary tree,
you need to output the second minimum value in the set made of all the nodes' value in the whole tree.

If no such second minimum value exists, output -1 instead.

Example 1:

Input:
    2
   / \
  2   5
     / \
    5   7

Output: 5
Explanation: The smallest value is 2, the second smallest value is 5.


Example 2:

Input:
    2
   / \
  2   2

Output: -1
Explanation: The smallest value is 2, but there isn't any second smallest value.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
// https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/discuss/371342/go-solution%3A-more-like-DFS
func findSecondMinimumValue(root *TreeNode) int {
	min := -1
	getMin(root, &min)

	return min
}

func getMin(root *TreeNode, min *int) {
	if root.Left == nil || root.Right == nil {
		return
	}

	if root.Left.Val != root.Val {
		if *min == -1 {
			*min = root.Left.Val
		} else if root.Left.Val < *min {
			*min = root.Left.Val
		}
	} else {
		getMin(root.Left, min)
	}

	if root.Right.Val != root.Val {
		if *min == -1 {
			*min = root.Right.Val
		} else if root.Right.Val < *min {
			*min = root.Right.Val
		}
	} else {
		getMin(root.Right, min)
	}
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/discuss/493135/go-clean-code-0ms-beats-100

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil || root.Left == nil {
		return -1
	}
	left, right := root.Left.Val, root.Right.Val
	if left == root.Val {
		left = findSecondMinimumValue(root.Left)
	}
	if right == root.Val {
		right = findSecondMinimumValue(root.Right)
	}
	if left == -1 || right == -1 {
		return int(math.Max(float64(left), float64(right)))
	}
	return int(math.Min(float64(left), float64(right)))
}
*/
