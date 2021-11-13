package main

/*
Given an n-ary tree, return the preorder traversal of its nodes' values.
Nary-Tree input serialization is represented in their level order traversal,
each group of children is separated by the null value (See examples).


Follow up:
Recursive solution is trivial, could you do it iteratively?

Example1:
	  1
	/ | \
   3  2  4
  / \
 5   6

Input: root = [1,null,3,2,4,null,5,6]
Output: [1,3,5,6,2,4]

Example2:
		 1
	/   /  \   \
   2   3    4   5
      / \   |  / \
	 6   7  8  9 10
		 |  |  |
		11 12 13
		 |
		14

Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
Output: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]


Constraints:
The height of the n-ary tree is less than or equal to 1000
The total number of nodes is between [0, 10^4]
*/

type Node struct {
	Val      int
	Children []*Node
}

// func preorder(root *Node) []int {
// 	res := []int{}
// 	if root == nil {
// 		return res
// 	}

// 	stack := []*Node{root}
// 	for len(stack) > 0 {
// 		r := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		res = append(res, r.Val)

// 		tmp := []*Node{}
// 		for _, v := range r.Children {
// 			tmp = append([]*Node{v}, tmp...)
// 		}
// 		stack = append(stack, tmp...)
// 	}
// 	return res
// }

func preorder(root *Node) []int {
	res := make([]int, 0)
	visit(root, &res)
	return res
}

func visit(root *Node, res *[]int) {
	if root == nil {
		return
	}

	// 節點值 append 回 res
	*res = append(*res, root.Val)

	// 尋遍 child 節點，如果 root.Children 為 nil 則不進迴圈
	for _, v := range root.Children {
		visit(v, res)
	}
}

func main() {

}

/*
[Recursive] solution:
- https://leetcode.com/problems/n-ary-tree-preorder-traversal/discuss/578363/go-clean-code-0ms-beats-100

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	ret := []int{root.Val}
	for _, child := range root.Children {
		ret = append(ret, preorder(child)...)
	}
	return ret
}
*/

/*
[Iterative] solution:
- https://leetcode.com/problems/n-ary-tree-preorder-traversal/discuss/545827/Go%3A-Iterative-solution

func preorder(root *Node) []int {
    res := []int{}
    if root == nil {
        return res
    }

    stack := []*Node{root}
    for len(stack) > 0 {
        r := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        res = append(res, r.Val)

        tmp := []*Node{}
        for _, v := range r.Children {
            tmp = append([]*Node{v}, tmp...)
        }
        stack = append(stack, tmp...)
    }
    return res
}
*/
