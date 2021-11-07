package main

/*
Given the root of a binary tree with unique values and the values of two different nodes of the tree x and y,

return true if the nodes corresponding to the values x and y in the tree are cousins, or false otherwise.

Two nodes of a binary tree are cousins if they have the same depth with different parents.

Note that in a binary tree, the root node is at the depth 0, and children of each depth k node are at the depth k + 1.

Example 1:
    1
   / \
  2   3
 /
4

Input: root = [1,2,3,4], x = 4, y = 3
Output: false

Example 2:
  1
 / \
2   3
 \   \
 4    5

Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
Output: true

Example 3:
  1
 / \
2   3
 \
  4

Input: root = [1,2,3,null,4], x = 2, y = 3
Output: false

Constraints:
- The number of nodes in the tree is in the range [2, 100].
- 1 <= Node.val <= 100
- Each node has a unique value.
- x != y
- x and y are exist in the tree.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type DNode struct {
	depth int
	node  *TreeNode
}

// https://leetcode.com/problems/cousins-in-binary-tree/discuss/240992/Golang-BFS
// 使用 BFS 探索所有節點值，同時計算深度
func isCousins(root *TreeNode, x int, y int) bool {
	var xdepth, ydepth int
	var xnode, ynode *TreeNode // xnode, ynode 分別代表左右子節點數

	q := []DNode{} // q 作為數節點，同步計算深度
	// 初始節點 q 為 root 節點
	q = append(q, DNode{
		depth: 0,
		node:  root,
	})
	for len(q) > 0 {
		dnode := q[0]
		q = q[1:]
		node, depth := dnode.node, dnode.depth // 迭代 node 為 dnode 的節點以及其對應的深度
		if (node.Left != nil && node.Left.Val == x) || (node.Right != nil && node.Right.Val == x) {
			xnode = node
			xdepth = depth + 1
		}
		if (node.Left != nil && node.Left.Val == y) || (node.Right != nil && node.Right.Val == y) {
			ynode = node
			ydepth = depth + 1
		}
		if node.Left != nil {
			q = append(q, DNode{depth + 1, node.Left})
		}
		if node.Right != nil {
			q = append(q, DNode{depth + 1, node.Right})
		}
	}
	return xdepth == ydepth && xnode != ynode
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/cousins-in-binary-tree/discuss/375809/Go-golang-DFS-solution

func isCousins(root *TreeNode, x int, y int) bool {
    xDepth, yDepth, xParent, yParent := 0, 0, 0, 0
    dfs(root, x, y, 0, 0, &xDepth, &yDepth, &xParent, &yParent)
    return xDepth == yDepth && xParent != yParent
}

// 使用 DFS 同步計算深度
func dfs(node *TreeNode, x, y, p, depth int, xDepth, yDepth, xParent, yParent *int) {
    if node == nil { return }
    if node.Val == x {
        *xDepth = depth
        *xParent = p
    }
    if node.Val == y {
        *yDepth = depth
        *yParent = p
    }
    dfs(node.Left, x, y, node.Val, depth + 1, xDepth, yDepth, xParent, yParent)
    dfs(node.Right, x, y, node.Val, depth + 1, xDepth, yDepth, xParent, yParent)
}
*/
