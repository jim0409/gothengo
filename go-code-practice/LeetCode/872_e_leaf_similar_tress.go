package main

/*
Consider all the leaves of a binary tree, from left to right order,
the values of those leaves form a leaf value sequence.

     3
    / \
   5   1
 / |   | \
6  2   9  8
  / \
 7   4

For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.

Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

      3						  3
   /    \					/   \
  5      1				   5     1
 / \    / \				  / \   / \
6   2  9   8			 6   7 4   2
   / \							  / \
  7   4							 9   8

Example 1:
Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
Output: true

Example 2:
Input: root1 = [1], root2 = [1]
Output: true

Example 3:
Input: root1 = [1], root2 = [2]
Output: false

Example 4:
Input: root1 = [1,2], root2 = [2,2]
Output: true

  1				 1
 / \			/ \
2   3		   3   2

Example 5:
Input: root1 = [1,2,3], root2 = [1,3,2]
Output: false

Constraints:
The number of nodes in each tree will be in the range [1, 200].
Both of the given trees will have values in the range [0, 200].
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 題目用意: 比較子節點的葉值是否一致
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// 透過dfs尋遍 s1 & s2
	s1, s2 := []int{}, []int{}
	dfs(root1, &s1)
	dfs(root2, &s2)

	// 如果 s1 與 s2 的長度不同回傳錯誤
	if len(s1) != len(s2) {
		return false
	}

	// 根據題意，s1及s2的各項值應當相同
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*s = append(*s, root.Val)
	}
	dfs(root.Left, s)
	dfs(root.Right, s)
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/leaf-similar-trees/discuss/152429/go-recursive

func dfs(node *TreeNode, leaf *([]int)) {
    if node == nil {
        return
    }
    if node.Left == nil && node.Right == nil {
        *leaf = append(*leaf, node.Val)
        return
    }
    dfs(node.Left, leaf)
    dfs(node.Right, leaf)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    var t1, t2 []int
    dfs(root1, &t1)
    dfs(root2, &t2)
    if len(t1) != len(t2) {
        return false
    }
    for i, val := range t1 {
        if t2[i] != val {
            return false
        }
    }
    return true
}
*/
