// 根據一顆樹的 前序遍歷(preorder traversal) 與 中序遍歷(Inorder traversal) 建構一顆二元樹
// - https://medium.com/@cashbooktw/construct-binary-tree-with-preorder-and-inorder-traversal-c50fb945f00f

/*
	透過preorder可以得知第一個數為根，但不知道接下來的數是`左子樹`或`右子樹``
	
	inorder可以判斷，在該數字前面的為`左子樹`反之則為`右子樹`

	從preorder開始拿數字
	假設第一個數字為r0，則透過r0來查找inorder裡面的r0前面一位數，並將該位數放在r0的左邊，反之則放在r0的右邊
	
	根據左子樹與右子樹的這些元素，可以另外在preorder中判斷，哪些屬於左子樹，哪些屬於右子樹
*/

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0{
		return nil
	}

	// preorder 的第一根節點即為根節點
	res := &TreeNode{
		Val : preorder[0],
	}

	if len(preorder) == 1{
		return res
	}

	// 依據 val 回傳val在 nums 內的位置
	idx := func(val int, nums []int) int{
		for i,v:=range nums{
			if v == val{
				return i
			}
		}
		// 如果都找不到，表示 preorder 與 inorder 中有數據不吻合 => 錯誤，回傳 -1 讓外面的 idx 報錯
		return -1
	}(res.Val, inorder)

	if idx == -1 {
		panic("data error")
	}

	// 在中序(InOrder)遍歷中，根節點前面的即為 跟節點的左子樹 ，後面的即為 跟節點的右子樹
	res.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	res.Right = buildTree(preorder[idx+1:], inorder[idx+1:])

	return res
}


func (t *TreeNode) PostOrderTraversal(){
	if t.Left != nil{
		t.Left.PostOrderTraversal()
	}
	if t.Right != nil{
		t.Right.PostOrderTraversal()
	}
	fmt.Printf("%d :", t.Val)
}

func main(){
	preorder := []int{0,1,3,4,2,5}
	inorder := []int{3,1,4,0,2,5}

	tree := buildTree(preorder, inorder)
	tree.PostOrderTraversal()

}