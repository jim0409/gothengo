package main

/*
Given a non-empty binary tree,
return the average value of the nodes on each level in the form of an array.

Example 1:
Input:
    3
   / \
  9  20
    /  \
   15   7

Output: [3, 14.5, 11]
Explanation:
The average value of nodes on level 0 is 3,
on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].

Note:
The range of node's value is in the range of 32-bit signed integer.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	values := [][]int{}
	collect(root, 0, &values)

	res := []float64{}
	for _, v := range values {
		length, sum := len(v), 0
		for _, vv := range v {
			sum += vv
		}
		res = append(res, float64(sum)/float64(length))
	}

	return res
}

func collect(root *TreeNode, level int, values *[][]int) {
	if root == nil {
		return
	}

	for len(*values) < level+1 {
		*values = append(*values, []int{})
	}

	(*values)[level] = append((*values)[level], root.Val)

	collect(root.Left, level+1, values)
	collect(root.Right, level+1, values)
}

func main() {

}

/*
[go-solution with both bfs & dfs] solution:
- https://leetcode.com/problems/average-of-levels-in-binary-tree/discuss/640768/go-bfs-and-dfs



*/

/*
solution:
- https://leetcode.com/problems/average-of-levels-in-binary-tree/discuss/486495/go-clean-code-4ms-beats-100

func averageOfLevels(root *TreeNode) (avg []float64) {
	if root == nil {
		return
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		sum, n := 0, len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		avg = append(avg, float64(sum)/float64(n))
	}
	return
}
*/
