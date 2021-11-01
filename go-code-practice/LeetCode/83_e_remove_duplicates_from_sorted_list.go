package main

import "fmt"

/*
Given a sorted linked list, delete all duplicates such that each element appear only once.


Example 1:

Input: 1->1->2
Output: 1->2

Example 2:

Input: 1->1->2->3->3
Output: 1->2->3
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}

// 	curNode := head
// 	for curNode.Next != nil {
// 		tmpNode := curNode

// 		if tmpNode.Val == tmpNode.Next.Val {
// 			fmt.Println(tmpNode.Val)
// 			tmpNode = tmpNode.Next
// 		}

// 		curNode.Next = tmpNode.Next
// 		if curNode.Next != nil {
// 			curNode = curNode.Next
// 			// 會有edge case，當curNode迭代到倒數第二位，會無法再往下判斷
// 			// e.g. []int{ 1, 1, 1 }
// 			// tmpNode=1,1 , curlNode=1,1 --> curNode = 1 ... 無法再往下迭代
// 		}
// 	}

// 	return head
// }

func deleteDuplicates(head *ListNode) *ListNode {
	curNode := head
	for {
		// 判斷邏輯閘
		if curNode == nil || curNode.Next == nil {
			break
		}

		// 當curNode當下的值等同於下一個節點的值，將curNode的下一個值往下迭代
		// 這邊不用擔心curNode.Next.Next，因為前面已經先判斷curNode.Next != nil
		if curNode.Val == curNode.Next.Val {
			curNode.Next = curNode.Next.Next
		} else {
			curNode = curNode.Next
		}
	}
	return head
}

func main() {
	inputArr1 := []int{1, 1, 2}
	dumpList(deleteDuplicates(createLists(inputArr1)))
	fmt.Println()

	inputArr2 := []int{1, 1, 2, 2, 3, 3}
	dumpList(deleteDuplicates(createLists(inputArr2)))
	fmt.Println()

	inputArr3 := []int{1, 1, 1}
	dumpList(deleteDuplicates(createLists(inputArr3)))
}

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

/*
solution:
- https://leetcode.com/problems/remove-duplicates-from-sorted-list/discuss/301625/go-4ms


func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	for {
		if p == nil || p.Next == nil {
			break
		}
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}

*/
