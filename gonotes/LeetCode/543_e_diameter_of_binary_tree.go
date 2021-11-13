package main

/*
Given a binary tree,
you need to compute the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path
between any two nodes in a tree.

This path may or may not pass through the root.


Example:
Given a binary tree
          1
         / \
        2   3
       / \
      4   5
Return 3, which is the length of the path [4,2,1,3] or [5,2,1,3].

Note:
The length of path between two nodes is represented
by the number of edges between them.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS 先深後廣演算法
// - https://alrightchiu.github.io/SecondRound/graph-depth-first-searchdfsshen-du-you-xian-sou-xun.html
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	visit(root, &res)
	return res
}

func visit(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	left := visit(root.Left, res)   // 找出左節點中最大深度(值)
	right := visit(root.Right, res) // 找出右節點中最大深度(值)

	*res = max(*res, left+right) // 看是否比`當下深度`還要更深，深度=`左邊深度+右邊深度`
	return 1 + max(left, right)  // 回傳`當下深度`=(左or右)最大深度+1(當下節點)

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
[DFS]solution
- https://leetcode.com/problems/diameter-of-binary-tree/discuss/649403/go-dfs

func diameterOfBinaryTree(root *TreeNode) int {
    res := 0
    visit(root, &res)
    return res
}

func visit(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }

    left := visit(root.Left, res)
    right := visit(root.Right, res)

    *res = max(*res, left + right)
    return 1 + max(left, right)

}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
*/
