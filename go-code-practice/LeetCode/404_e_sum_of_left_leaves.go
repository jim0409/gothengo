package main

/*
Find the sum of all left leaves in a given binary tree.

Example:

    3
   / \
  9  20
    /  \
   15   7

There are two left leaves in the binary tree,
with values 9 and 15 respectively. Return 24.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftSum := sumOfLeftLeaves(root.Left)
	rightSum := sumOfLeftLeaves(root.Right)

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftSum += root.Left.Val
	}

	return leftSum + rightSum
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/sum-of-left-leaves/discuss/466017/Go-0ms-solution-very-simple


func sumOfLeftLeaves(root *TreeNode) int {
    return helper(root, false)
}

func helper(node *TreeNode, isLeft bool) int {
    if node == nil {
        return 0
    }
    if node.Left == nil && node.Right == nil {
        if isLeft {
            return node.Val
        }
        return 0
    }
    v1 := helper(node.Left, true)
    v2 := helper(node.Right, false)
    return v1+v2
}


solution2:
- https://leetcode.com/problems/sum-of-left-leaves/discuss/340802/Go-O(n)-with-stack


func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	stack := []*TreeNode{}

	for root != nil || len(stack) != 0 {
		for root != nil {
			if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
				sum += root.Left.Val
			}
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		root = node.Right
	}

	return sum
}
*/
