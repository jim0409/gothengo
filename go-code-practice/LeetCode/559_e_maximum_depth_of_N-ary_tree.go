package main

/*
Given a n-ary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path
from the root node down to the farthest leaf node.

Nary-Tree input serialization is represented in their level order traversal,
each group of children is separated by the null value (See examples).

Example1:
Input: root = [1,null,3,2,4,null,5,6]
		1
	  / |  \
	3	2	4
   / \
  5   6

Output: 3


Example2:
Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: 5


Constraints:
The depth of the n-ary tree is less than or equal to 1000.
The total number of nodes is between [0, 10^4].
*/

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	m := 1
	for _, v := range root.Children {
		m = max(m, maxDepth(v)+1)
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}

/*
solution
- https://leetcode.com/problems/maximum-depth-of-n-ary-tree/discuss/578285/go-clean-code-0ms-beats-100

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var depth int
	for _, child := range root.Children {
		depth = int(math.Max(float64(depth), float64(maxDepth(child))))
	}
	return 1 + depth
}
*/
