package main

import "math"

/*
Given a binary tree, return the tilt of the whole tree.

The tilt of a tree node is defined as

`the absolute difference between the sum of all left subtree node values`
and the sum of all right subtree node values.`

Null node has tilt 0.

The tilt of the whole tree is defined as the sum of all nodes' tilt.

Example:
Input:
         1
       /   \
      2     3
Output: 1
Explanation:
Tilt of node 2 : 0
Tilt of node 3 : 0
Tilt of node 1 : |2-3| = 1
Tilt of binary tree : 0 + 0 + 1 = 1


Note:

The sum of node values in any subtree won't exceed the range of 32-bit integer.
All the tilt values won't exceed the range of 32-bit integer.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// tilt: 傾斜
// 一個BTree的傾斜 = 該treeNode所有傾斜的和 = 左節點傾斜和 + 右節點傾斜和
// https://leetcode.com/problems/binary-tree-tilt/discuss/388316/Go-golang-clean-solution
func findTilt(root *TreeNode) int {
	ans := 0
	helper(root, &ans)
	return ans
}

func helper(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left, ans)                        // 計算左節點傾斜
	right := helper(root.Right, ans)                      // 計算右節點傾斜
	*ans += int(math.Abs(float64(left) - float64(right))) // 計算當前節點左 + 右 傾斜
	return root.Val + left + right
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/binary-tree-tilt/discuss/572560/Go%3A-bottom-up-solution


func findTilt(root *TreeNode) int {
    _, res := tilt(root)
    return res
}

func tilt(root *TreeNode) (int,int) {
    if root == nil {
        return 0, 0
    }
    left, s1 := tilt(root.Left)
    right, s2 := tilt(root.Right)
    return (root.Val+left+right), (abs(left-right)+s1+s2)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
*/
