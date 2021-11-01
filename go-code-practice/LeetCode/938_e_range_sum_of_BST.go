package main

/*
Given the root node of a binary search tree and two integers low and high,
return `the sum of values of all nodes` with a value in the inclusive range [low, high].


Example 1:
	 10
	/  \
   5    15
  / \     \
 3   7     18
Input: root = [10, 5, 15, 3, 7, null, 18], low = 7, high = 15
Output: 32
Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.


Example 2:
	   10
	 /   \
    5     15
   / \    /  \
  3   7  13  18
 /   /
1   6
Input: root = [10, 5, 15, 3, 7, 13, 18, 1, null, 6], low = 6, high = 10
Output: 23
Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.


Constraints:
The number of nodes in the tree is in the range [1, 2 * 104].
1 <= Node.val <= 105
1 <= low <= high <= 105
All Node.val are unique.
*/

/*
題意:
給定一個二元樹，以及一個範圍值
將樹中所有節點值座落於該範圍的值加總起來後回傳

1. 尋遍所有數的節點值
2. 判斷節點值是否符合範圍值條件，符合則進行加總。反之則跳過
3. 回傳最後加總值
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/range-sum-of-bst/discuss/637185/go-recursive
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	res := 0

	// 確認 Node 節點的值是否座落於範圍內，是的話將其加入至res
	if root.Val >= L && root.Val <= R {
		res += root.Val
	}

	// 判斷，如果當節點的值大於最大值。搜尋左節點的值(通常小於該節點)。做遞迴
	if root.Val > L {
		res += rangeSumBST(root.Left, L, R)
	}

	// 判斷，如果當節點的值小於最小值。搜尋右節點的值(通常大於該節點)。做遞迴
	if root.Val < R {
		res += rangeSumBST(root.Right, L, R)
	}

	return res
}

func main() {}

/*
refer:
- https://leetcode.com/problems/range-sum-of-bst/discuss/719855/Go-Level-Traversal

func rangeSumBST(root *TreeNode, L int, R int) int {
	var result int
	if root != nil {
		q := []*TreeNode{root}

		for len(q) > 0 {
			n := q[0]
			q = q[1:]
			if n.Val >= L && n.Val <= R {
				result += n.Val
			}

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
	}

	return result
}


- https://leetcode.com/problems/range-sum-of-bst/discuss/523627/Go-golang-BFS
func rangeSumBST(root *TreeNode, L int, R int) int {
    stack := []*TreeNode{root}
    sum := 0
    p := 0
    for p < len(stack) {
        if stack[p].Val >= L && stack[p].Val <= R {
            sum += stack[p].Val
            if stack[p].Left != nil {
                stack = append(stack, stack[p].Left)
            }
            if stack[p].Right != nil {
                stack = append(stack, stack[p].Right)
            }
        } else if stack[p].Val < L {
            if stack[p].Right != nil {
                stack = append(stack, stack[p].Right)
            }
        } else if stack[p].Val > R {
            if stack[p].Left != nil {
                stack = append(stack, stack[p].Left)
            }
        }
        p ++
    }
    return sum
}

- https://leetcode.com/problems/range-sum-of-bst/discuss/429417/Go-golang-DFS-solution
func rangeSumBST(root *TreeNode, low int, high int) int {
    ans := 0
    dfs(root, low, high, &ans)
    return ans
}

func dfs(root *TreeNode, low, high int, ans *int) {
    if root == nil { return }
    if root.Val >= low && root.Val <= high { *ans += root.Val }
    if root.Val >= low { dfs(root.Left, low, high, ans) }
    if root.Val <= high { dfs(root.Right, low, high, ans) }

}
*/
