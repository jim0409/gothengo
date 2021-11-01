package main

/*
Given a Binary Search Tree (BST) with the root node root,
return the minimum difference between the values of any two different nodes in the tree.

Example :
Input: root = [4,2,6,1,3,null,null]
Output: 1
Explanation:
Note that root is a TreeNode object, not an array.

The given tree [4,2,6,1,3,null,null] is represented by the following diagram:

          4
        /   \
      2      6
     / \
    1   3

while the minimum difference in this tree is 1, it occurs between node 1 and node 2, also between node 3 and node 2.
Note:

The size of the BST will be between 2 and 100.
The BST is always valid, each node's value is an integer, and each node's value is different.
This question is the same as 530: https://leetcode.com/problems/minimum-absolute-difference-in-bst/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	res := 1 << 8
	c := []int{}

	bsf(root, &c) // BSF 先把節點攤平

	for i := 1; i < len(c); i++ {
		res = min(res, c[i]-c[i-1]) // 將每個值分別取出，比較出節點前後差異最小值
	}

	return res
}

func bsf(root *TreeNode, c *[]int) {
	if root == nil {
		return
	}

	bsf(root.Left, c)
	*c = append(*c, root.Val)

	bsf(root.Right, c)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/minimum-distance-between-bst-nodes/discuss/724239/Go-Iterative-solution-without-recursion

func minDiffInBST(root *TreeNode) int {
	st := &Stack{}
	st.push(root)
	prev := (1 << 31) * -1
	min := 1 << 31
	visited := map[*TreeNode]struct{}{}
	for st.len() > 0 {
		el := st.pop()
		if el == nil {
			continue
		}
		if _, ok := visited[el]; ok {
			if el.Val-prev < min {
				min = el.Val - prev
			}
			prev = el.Val
		} else {
			visited[el] = struct{}{}
			st.push(el.Right)
			st.push(el)
			st.push(el.Left)
		}
	}
	return min
}

type Stack struct {
	arr []*TreeNode
}

func (s *Stack) push(i *TreeNode) {
	s.arr = append(s.arr, i)
}

func (s *Stack) pushAll(i []*TreeNode) {
	s.arr = append(s.arr, i...)
}

func (s *Stack) pop() *TreeNode {
	el := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return el
}

func (s *Stack) front() *TreeNode {
	el := s.arr[0]
	s.arr = s.arr[1:len(s.arr)]
	return el
}

func (s *Stack) len() int {
	return len(s.arr)
}
*/
