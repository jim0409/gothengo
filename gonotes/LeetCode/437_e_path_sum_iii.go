package main

/*
You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

計算排列組合。在給定的btree下，算出符合總和值為8的路徑的可能情境

Example:

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

Return 3. The paths that sum to 8 are:

1.  5 -> 3
2.  5 -> 2 -> 1
3. -3 -> 11
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return helper(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

// helper: 節點下達到期望值的可能性，左/右節點，分別當作root節點後去迭代出可能性
func helper(root *TreeNode, sum int) int {
	count := 0
	if root == nil {
		return 0
	}
	if root.Val == sum {
		count++
	}
	return count + helper(root.Left, sum-root.Val) + helper(root.Right, sum-root.Val)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/path-sum-iii/discuss/299928/Go-Golang-easy-to-follow-4ms-100


func pathSum(root *TreeNode, sum int) int {
    m := make(map[int]int, 0)
    m[0] = 1
    return helper(root, 0, sum, m)
}

func helper(root *TreeNode, curSum, sum int, m map[int]int) int{
    if root == nil {
        return 0
    }
    curSum += root.Val
    total := 0
    if val, ok := m[curSum - sum]; ok {
        total = val
    }
    m[curSum] = m[curSum] + 1
    total += helper(root.Left, curSum, sum , m) + helper(root.Right, curSum, sum , m)
    m[curSum] = m[curSum] -1

    return total
}
*/

/*
solution2:
- https://leetcode.com/problems/path-sum-iii/discuss/398148/Go-golang-clean-solution


func pathSum(root *TreeNode, sum int) int {
    if root == nil { return 0 }
    return helper(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func helper(root *TreeNode, sum int) int {
    count := 0
    if root == nil { return 0 }
    if root.Val == sum { count++ }
    return count + helper(root.Left, sum - root.Val) + helper(root.Right, sum - root.Val)
}
*/
