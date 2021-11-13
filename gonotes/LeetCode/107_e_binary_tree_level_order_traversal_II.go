package main

import "fmt"

/*
Given a binary tree,

return the bottom-up level order traversal of its nodes' values.
(ie, from left to right, level by level from leaf to root).


For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7

return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
*/

func levelOrderBottom(root *TreeNode) [][]int {
	// levelOrderUp，從上到下組織出[ [root.Val], [root.Right.Val, root.Left.Val], ... ]
	res := levelOrderUp(root)
	// 將原本的陣列反過來排列
	reverseSlice(res)
	return res
}

func levelOrderUp(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// 定義出接下來要append的陣列
	var res [][]int
	var q []*TreeNode
	q = append(q, root)

	// 初始化的 q = []*TreeNode{ root }, len(q) = 1
	// q 是一個動態的 []*TreeNode 陣列
	for len(q) > 0 {

		// l 定義為len(q), 評估有幾個節點需要進行計算
		// 每次在迭代len(q)的時候，都會初始化 `tmp := []int{}``
		l := len(q)
		var tmp []int

		// l 為len(q)，因為 q = append(q[1:], node.Left, node.Right )
		// 所以 l 可能會因為 q 這個陣列增加而增加
		for i := 0; i < l; i++ {
			// 把node拉出來
			node := q[0]

			// [第一次] 剔除掉 root 的節點 ... 3
			// [第二次] 剔除掉 root.Left 的節點 ... 9
			// [第三次] 剔除掉 root.Right 的節點 ... 2
			q = q[1:]

			// 觀察node下左右兩邊的節點，看節點是否非空，非空就把節點append進q
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			// l 為 len(q)
			// [第一次] l = 1
			// [第二次] l = 2
			tmp = append(tmp, node.Val)
		}

		// tmp = append(tmp, node.Val, node.Left.Val, node.Right.Val ... ) => []int{ vRv, vLv, vRv, ...}
		res = append(res, tmp)
	}
	return res
}

func reverseSlice(s [][]int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func main() {
	intArr := []int{3, 9, 2}
	tree := ImplementTree(intArr)
	fmt.Println(levelOrderBottom(tree))

	intArr1 := []int{3, 9, 2, 8, 7, 1, 0}
	tree1 := ImplementTree(intArr1)
	fmt.Println(levelOrderBottom(tree1))
}

/*
[GO BFS] solution:
- https://leetcode.com/problems/binary-tree-level-order-traversal-ii/discuss/378548/go-bfs


func levelOrderBottom(root *TreeNode) [][]int {
	res := levelOrderUp(root)
	reverseSlice(res)
	return res
}

func levelOrderUp(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		l := len(q)
		var tmp []int
		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
	}
	return res
}

func reverseSlice(s [][]int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ImplementTree(arr []int) *TreeNode {
	rtNode := TreeNode{arr[0], nil, nil}

	for i := 1; i < len(arr); i++ {
		rtNode.InsertNode(arr[i])
	}

	return &rtNode
}

func (n *TreeNode) InsertNode(v int) {

	if n == nil {
		n = &TreeNode{v, nil, nil}

	} else {

		if v < n.Val {
			if n.Left == nil {
				n.Left = &TreeNode{v, nil, nil}
			} else {
				n.Left.InsertNode(v)
			}
		} else {
			if n.Right == nil {
				n.Right = &TreeNode{v, nil, nil}
			} else {
				n.Right.InsertNode(v)
			}
		}
	}
}
