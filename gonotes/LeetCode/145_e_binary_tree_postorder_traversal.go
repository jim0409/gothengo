package main

/*
Given the root of a binary tree, return the postorder traversal of its nodes' values.

Example 1:
1
 \
  2
 /
3
Input: root = [1,null,2,3]
Output: [3,2,1]

Example 2:
Input: root = []
Output: []

Example 3:
Input: root = [1]
Output: [1]

Example 4:
  1
 /
2
Input: root = [1,2]
Output: [2,1]

Example 5:
1
 \
  2
Input: root = [1,null,2]
Output: [2,1]


Constraints:

The number of the nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100

Follow up: Recursive solution is trivial, could you do it iteratively?
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	postOrder := make([]int, 0)
	postOrder = append(postOrder, postorderTraversal(root.Left)...)
	postOrder = append(postOrder, postorderTraversal(root.Right)...)
	postOrder = append(postOrder, root.Val)

	return postOrder
}

func main() {}

/*
refer:
- https://leetcode.com/problems/binary-tree-postorder-traversal/discuss/964196/go-clean-code-0ms-beats-100

func postorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var ret []int
	for root != nil || len(stack) > 0 {
		for ; root != nil; root = root.Right {
			ret = append([]int{root.Val}, ret...)
			stack = append(stack, root)
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Left
	}
	return ret
}



- https://leetcode.com/problems/binary-tree-postorder-traversal/discuss/244798/Go-iterative-solution-with-Stack

func postorderTraversal(root *TreeNode) []int {
    var result []int
    var s Stack
    if root == nil {
        return result
    }
    s.Push(*root)
    for !s.IsEmpty() {
        temp := s.Pop()
        if temp.Left != nil {
            s.Push(*(temp.Left))
        }
        if temp.Right != nil {
            s.Push(*(temp.Right))
        }
        result = append(result, temp.Val)
    }

    for i, j := 0, len(result) - 1; i < len(result) / 2; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    return result
}

type Stack []TreeNode

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}

func (s *Stack) Push(node TreeNode) {
    *s = append(*s, node)
}

func (s *Stack) Pop() TreeNode {
    node := (*s)[len(*s) - 1]
    *s = (*s)[0 : len(*s) - 1]
    return node
}


- https://leetcode.com/problems/binary-tree-postorder-traversal/discuss/417203/go-concise-stack-100

func postorderTraversal(root *TreeNode) []int {
	// 假設 root 為空，則回傳空陣列
	if root == nil {
		return nil
	}

	// 宣告一個樹節點，用來做後面建樹用
	type trackTreeNode struct {
		root    *TreeNode
		visited bool
	}

	// 宣告一個樹節點的stack方便後面紀錄 postorder 用
	var (
		stack []*trackTreeNode
		res   []int
	)

	// 先綁定第一個 root 節點
	stack = append(stack, &trackTreeNode{root: root})

	// 當stack長度不為 0 時一直迭代下去
	for len(stack) != 0 {
		node := stack[len(stack)-1] // peek
		if node.visited {
			stack = stack[:len(stack)-1] // pop
			res = append(res, node.root.Val)
			continue
		}
		if node.root.Right != nil {
			stack = append(stack, &trackTreeNode{root: node.root.Right}) // push right child
		}
		if node.root.Left != nil {
			stack = append(stack, &trackTreeNode{root: node.root.Left}) // push left child
		}
		node.visited = true
	}

	return res
}
*/
