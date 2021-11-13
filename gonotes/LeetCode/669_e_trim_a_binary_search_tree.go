package main

/*
Given a binary search tree and the lowest and highest boundaries as L and R,
trim the tree so that all its elements lies in [L, R] (R >= L).

You might need to change the root of the tree,
so the result should return the new root of the trimmed binary search tree.


Example 1:
Input:
    1
   / \
  0   2

  L = 1
  R = 2

Output:
    1
      \
       2


Example 2:
Input:
    3
   / \
  0   4
   \
    2
   /
  1

  L = 1
  R = 3

Output:
      3
     /
   2
  /
 1
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  trimBST(root.Left, L, R),
		Right: trimBST(root.Right, L, R),
	}
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/trim-a-binary-search-tree/discuss/493266/Go-100-solution

func trimBST(root *TreeNode, L int, R int) *TreeNode {
    return walk(root, L, R)
}

func walk(node *TreeNode, L, R int) *TreeNode {
    if node == nil {
        return nil
    }
    if node.Val < L {
        return walk(node.Right, L, R)
    } else if node.Val > R {
        return walk(node.Left, L, R)
    } else {
        node.Left = walk(node.Left, L, R)
        node.Right = walk(node.Right, L, R)
        return node
    }
}
*/
