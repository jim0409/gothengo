package main

import "strconv"

/*
Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:

Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]

Explanation: All root-to-leaf paths are: 1->2->5, 1->3
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func binaryTreePaths(root *TreeNode) []string {
// 	if root == nil {
// 		return nil
// 	}
// 	path := []string{}
// 	path = append(path, fmt.Sprintf("%d", root.Val))
// 	path = append(path, binaryTreePaths(root.Left)...)
// 	path = append(path, binaryTreePaths(root.Right)...)

// 	return path
// }

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	result := []string{}
	helper(root, strconv.Itoa(root.Val), &result)
	return result
}

func helper(root *TreeNode, str string, result *[]string) {
	if root.Left == nil && root.Right == nil {
		*result = append(*result, str)
		return
	}
	if root.Left != nil {
		helper(root.Left, str+"->"+strconv.Itoa(root.Left.Val), result)
	}
	if root.Right != nil {
		helper(root.Right, str+"->"+strconv.Itoa(root.Right.Val), result)
	}
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/binary-tree-paths/discuss/235583/Go-0ms


func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return []string{}
    }

    return binaryTreePathsAccumulator(root, "")

}

func binaryTreePathsAccumulator(root *TreeNode, currentPath string) []string {

    if currentPath != ""{
        currentPath = currentPath+"->"+ strconv.Itoa(root.Val)
    }else {
        currentPath = strconv.Itoa(root.Val)
    }

    if root.Left==nil && root.Right == nil {
        return []string{currentPath}
    }

    result := []string{}
    if root.Left != nil {
        result = binaryTreePathsAccumulator(root.Left, currentPath)
    }

    if root.Right != nil {
        result = append(result, binaryTreePathsAccumulator(root.Right, currentPath)...)
    }

    return result
}
*/
