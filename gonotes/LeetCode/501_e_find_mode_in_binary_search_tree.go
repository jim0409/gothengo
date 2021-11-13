package main

/*
Given a binary search tree (BST) with duplicates,
find all the mode(s) (the most frequently occurred element) in the given BST.

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys `less than or equal to the node's key`.
The right subtree of a node contains only nodes with keys `greater than or equal to the node's key`.
Both the left and right subtrees must also be binary search trees.


For example:
Given BST [1,null,2,2],

   1
    \
     2
    /
   2


return [2].

Note: If a tree has more than one mode, you can return them in any order.

Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/find-mode-in-binary-search-tree/discuss/595319/Go-iterative-approach-summing-the-nodes
func findMode(root *TreeNode) []int {
	// brute force : unroll the tree and make a map of the entries
	// then sum them up
	// 暴力法：把所有的樹節點都攤平後，用一個map把裡面的值都存下來
	m := map[int]int{}
	q := []*TreeNode{}
	if root == nil {
		return nil
	}

	// unroll the tree and map the elements
	q = append(q, root)
	for len(q) > 0 {
		t := q[0]
		if t.Left != nil {
			q = append(q, t.Left)
		}
		if t.Right != nil {
			q = append(q, t.Right)
		}
		m[t.Val] += 1
		q = q[1:]
		// q = []TreeNode{ root, root.Left, root.Right, ...}
	}

	// Now sum up the entries and create a return slice of the max vals
	max := 0
	ret := []int{}
	for k, v := range m {
		// 如果 v > max, 將原本的ret重置為空,並將 `新的v` append 進 ret
		if v > max {
			ret = nil
			ret = append(ret, k)
			max = v
		} else if v == max {
			ret = append(ret, k)
		}
	}
	return ret
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/find-mode-in-binary-search-tree/discuss/468703/go-clean-code-4ms-beats-100


func findMode(root *TreeNode) []int {
	var mode []int
   var stack []*TreeNode

	for max, cnt, prev := 0, 0, 0; root != nil || len(stack) > 0; root = root.Right {

		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val == prev {
			cnt++
		} else {
			cnt, prev = 1, root.Val
      }

		if cnt > max {
			max = cnt
			mode = []int{prev}
		} else if cnt == max {
			mode = append(mode, prev)
		}
	}
	return mode
}
*/
