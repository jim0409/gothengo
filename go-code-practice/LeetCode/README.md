# 代碼解題

# refer:
### origin(EASY)
- https://leetcode.com/problemset/all/?difficulty=Easy

### solution refernce
- https://skyyen999.gitbooks.io/-leetcode-with-javascript/content/questions/28md.html


### 常用函數
```go
// create link and list
func createLists(arr []int) *ListNode {
	head := &ListNode{Val: arr[0], Next: nil}
	curNode := head
	for i := 1; i < len(arr); i++ {
		tmpNode := &ListNode{Val: arr[i], Next: nil}
		curNode.Next = tmpNode
		curNode = curNode.Next
	}
	return head
}

func dumpList(l *ListNode) {
	for l != nil {
		fmt.Printf("%v ", l.Val)
		l = l.Next
	}
}
```

```go
// create binary tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ImplementTree(arr []int) *TreeNode {
	rtNode := TreeNode{arr[0], nil, nil}

	for i := 1; i < len(arr); i++ {
		rtNode.InsertNode(arr[i])
	}

	return &rtNode
}

func (n *TreeNode) InsertNode(v int) {

	if n == nil {
		n = &TreeNode{v, nil, nil}

	} else {

		if v < n.Val {
			if n.Left == nil {
				n.Left = &TreeNode{v, nil, nil}
			} else {
				n.Left.InsertNode(v)
			}
		} else {
			if n.Right == nil {
				n.Right = &TreeNode{v, nil, nil}
			} else {
				n.Right.InsertNode(v)
			}
		}
	}
}

```

# refer:
為什麼要用`#!bin/bash`
- https://kknews.cc/zh-tw/code/96vkgol.html