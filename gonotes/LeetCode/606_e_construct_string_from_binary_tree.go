package main

import (
	"strconv"
	"strings"
)

/*
You need to construct a string consists of parenthesis and integers from a binary tree
with the preorder traversing way.

The null node needs to be represented by empty parenthesis pair "()".
And you need to omit all the empty parenthesis pairs that
don't affect the one-to-one mapping relationship between the string and the original binary tree.

Example 1:
Input: Binary tree: [1,2,3,4]
       1
     /   \
    2     3
   /
  4

Output: "1(2(4))(3)"

Explanation:
Originallay it needs to be "1(2(4)())(3()())",
but you need to omit all the unnecessary empty parenthesis pairs.
And it will be "1(2(4))(3)".

Example 2:
Input: Binary tree: [1,2,3,null,4]
       1
     /   \
    2     3
     \
      4

Output: "1(2()(4))(3)"

Explanation:
Almost the same as the first example,
except we can't omit the first parenthesis pair
to break the one-to-one mapping relationship between the input and the output.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/construct-string-from-binary-tree/discuss/303974/go
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	var str []string
	preorderStr(t, &str)

	s := strings.Join(str, "")
	s = strings.TrimPrefix(s, "(")
	s = strings.TrimSuffix(s, ")")
	return s
}

func preorderStr(t *TreeNode, str *[]string) {
	*str = append(*str, "(", strconv.Itoa(t.Val))

	if t.Left == nil && t.Right == nil {
		*str = append(*str, ")")
		return
	}

	if t.Left == nil {
		*str = append(*str, "()")
	} else {
		preorderStr(t.Left, str)
	}

	if t.Right != nil {
		preorderStr(t.Right, str)
	}
	*str = append(*str, ")")
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/construct-string-from-binary-tree/discuss/518675/Go-4ms-100-solution

func tree2str(t *TreeNode) string {
        if t == nil {
            return ""
        }
        if t.Right == nil && t.Left == nil {
            return strconv.Itoa(t.Val)
        } else if t.Right == nil{
            return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"
        } else {
            return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")" + "(" + tree2str(t.Right) + ")"
        }
}
*/
