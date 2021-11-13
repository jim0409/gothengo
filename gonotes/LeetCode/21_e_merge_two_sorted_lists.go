package main

import "fmt"

// Merge two sorted linked lists
// and return it as a new list.

// The new list should be made by splicing
// together the nodes of the first two lists.

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 製作一個空頭的header
	var newList = &ListNode{}
	// 建立一個動態的指標，並且把空頭header的位址也賦予給他
	var out = newList

	// 判斷 l1 & l2 皆非空才往下做
	// 因為如果 l1 為空，則 newList 可以直接 appned l2，反之亦然
	for l1 != nil && l2 != nil {
		// 如果 l1.Val 比 l2.Val 小，就拿 l1 的值並且將l1像下迭代
		if l1.Val < l2.Val {
			newList.Next = l1
			l1 = l1.Next
			newList = newList.Next
		} else {
			newList.Next = l2
			l2 = l2.Next
			newList = newList.Next
		}
	}

	// 跳出迴圈，表示 l1 或 l2 其中之一為空，將剩餘數列append上去
	if l1 != nil {
		newList.Next = l1
	} else if l2 != nil {
		newList.Next = l2
	}

	return out.Next
}

func main() {
	// Input: 1->2->4, 1->3->4
	// Output: 1->1->2->3->4->4
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	arL1 := buidListNode(arr1)
	arL2 := buidListNode(arr2)
	mergeL := mergeTwoLists(arL1, arL2)
	mergeL.print()

}

func (l *ListNode) print() {
	for l.Next != nil {
		fmt.Printf("%v ", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func buidListNode(arr []int) *ListNode {
	head := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	curNode := head
	for i := 1; i < len(arr); i++ {
		tmpNode := &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		curNode.Next = tmpNode
		curNode = curNode.Next
	}

	return head
}
