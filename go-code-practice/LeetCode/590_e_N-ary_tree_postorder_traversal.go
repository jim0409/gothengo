package main

/*
Given an n-ary tree, return the postorder traversal of its nodes' values.
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
Output: [5,6,3,2,4,1]

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
Output: [2,6,14,11,7,3,12,8,4,13,9,10,5,1]


Constraints:
The height of the n-ary tree is less than or equal to 1000
The total number of nodes is between [0, 10^4]
*/

type Node struct {
	Val      int
	Children []*Node
}

// https://leetcode.com/problems/n-ary-tree-postorder-traversal/discuss/637178/go-recursive
func postorder(root *Node) []int {
	res := make([]int, 0)
	visit(root, &res)
	return res
}

func visit(root *Node, res *[]int) {
	if root == nil {
		return
	}

	// 尋遍 child 節點，如果 root.Children 為 nil 則不進迴圈
	for _, v := range root.Children {
		visit(v, res)
	}

	// 節點值 append 回 res
	*res = append(*res, root.Val)
}

func main() {

}

/*
[Recursive] solution:
- https://leetcode.com/problems/n-ary-tree-postorder-traversal/discuss/578374/go-clean-code-0ms-beats-100

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var ret []int
	for _, child := range root.Children {
		ret = append(ret, postorder(child)...)
	}
	ret = append(ret, root.Val)
	return ret
}
*/

/*
[Iterative] solution:
- https://leetcode.com/problems/n-ary-tree-postorder-traversal/discuss/549341/Go%3A-Iterative-solution-(using-list)

func postorder(root *Node) []int {
    res := []int{}
    if root == nil {
        return res
    }

    stack := list.New()
    stack.PushFront(root)
    for stack.Len() > 0 {
        curr := stack.Remove(stack.Front()).(*Node)
        if curr != nil {
            res = append(res, curr.Val)
        }

        for i := 0; i < len(curr.Children) ; i++ {
            stack.PushFront(curr.Children[i])
        }
    }

    reverse(res)

    return res
}

func reverse (s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}
*/
