package main

/*
Given a binary search tree, rearrange the tree in in-order
so that the leftmost node in the tree is now the root of the tree,
and every node has no left child and only 1 right child.


Example 1:
Input: [5,3,6,2,4,null,8,1,null,null,null,7,9]

       5
      / \
    3    6
   / \    \
  2   4    8
 /        / \
1        7   9

Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

 1
  \
   2
    \
     3
      \
       4
        \
         5
          \
           6
            \
             7
              \
               8
                \
                 9


Constraints:
The number of nodes in the given tree will be between 1 and 100.
Each node will have a unique integer value from 0 to 1000.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	min, _ := helper(root)
	return min
}

func helper(node *TreeNode) (*TreeNode, *TreeNode) {
	min, max := node, node

	// 先找左邊節點，將左邊節點換成最小值。將節點右邊第一個子節點換成節點本身
	if node.Left != nil {
		lMin, lMax := helper(node.Left)
		min = lMin
		lMax.Right = node
	}

	// 比較右節點，將節點右邊換成 rMin
	if node.Right != nil {
		rMin, rMax := helper(node.Right)
		max = rMax
		node.Right = rMin
	}

	node.Left = nil
	return min, max
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/increasing-order-search-tree/discuss/581742/go-clean-code-0ms-beats-100


func increasingBST(root *TreeNode) *TreeNode {
	ret := &TreeNode{}
	prev := ret
	var stack []*TreeNode
	for ; root != nil || len(stack) != 0; root = root.Right {
		for ; root != nil; root = root.Left {
			stack = append(stack, root)
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		prev.Left, prev.Right = nil, root
		prev = prev.Right
	}
	return ret.Right
}
*/
