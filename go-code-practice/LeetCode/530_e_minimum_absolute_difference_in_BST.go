package main

import "math"

/*
Given a binary search tree with non-negative values,
find the minimum absolute difference between values of any two nodes.

Example:

Input:

   1
    \
     3
    /
   2

Output:
1

Explanation:
The minimum absolute difference is 1, which is the difference between 2 and 1
(or between 2 and 3).


Note:

There are at least two nodes in this BST.
This question is the same as 783:
- https://leetcode.com/problems/minimum-distance-between-bst-nodes/

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/discuss/502229/Go-100-solution
func getMinimumDifference(root *TreeNode) int {
	_, _, diff := helper(root)
	return diff
}

func helper(node *TreeNode) (int, int, int) {
	min, max, diff := node.Val, node.Val, math.MaxInt32

	if node.Left != nil {
		// 計算左邊節點的 helper
		lMin, lMax, lDiff := helper(node.Left)
		min = lMin // return min
		if tempDiff := node.Val - lMax; tempDiff < diff {
			diff = tempDiff
		}
		if lDiff < diff {
			diff = lDiff
		}
	}

	if node.Right != nil {
		// 計算右邊節點的 helper
		rMin, rMax, rDiff := helper(node.Right)
		max = rMax // return max
		if tempDiff := rMin - node.Val; tempDiff < diff {
			diff = tempDiff
		}
		if rDiff < diff {
			diff = rDiff
		}
	}

	return min, max, diff
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/minimum-absolute-difference-in-bst/discuss/468983/go-clean-code-8ms-beats-100

func getMinimumDifference(root *TreeNode) int {
	min, prev := math.MaxInt32, math.MinInt32
	var stack []*TreeNode
	for ; root != nil || len(stack) != 0; root = root.Right {
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		min = int(math.Min(float64(min), float64(root.Val-prev)))
		prev = root.Val
	}
	return min
}
*/
