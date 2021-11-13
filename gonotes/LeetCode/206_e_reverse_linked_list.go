package main

import "fmt"

/*
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively.
Could you implement both?
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 如果 head 為 nil 就直接返回
	if head == nil {
		return head
	}

	/*
		dummy: 空node
		pre: 用來置換 current node
		cur: 紀錄 head 迭代下來的 node 節點
	*/
	dummy := &ListNode{}
	
	// 初始化下 dummy 跟 curl 都是 head 值
	dummy.Next = head
	cur := dummy.Next
	var pre *ListNode

	// 利用 cur 進行迭代
	for cur != nil {
		/*
			next為每次生成的 tmp Node
			cur預設為dumm.Next=head，此處迭代。每次為cur.Next會指向pre(前一次的cur)
			pre = cur (讓下次cur在指的時候可以找到前一次的cur)
			cur = next(定義cur為tmp Node)
		*/

		next := cur.Next
		// 將cur的下一個節點指向pre，但是pre為前一次的cur
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}


func main() {
	listArr:=[]int{1,2,3,4}
	listNode1 := createLists(listArr)
	reverseList(listNode1)
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

/*
solution:
- https://leetcode.com/problems/reverse-linked-list/discuss/537014/Go


func reverseList(head *ListNode) *ListNode {
    var front *ListNode
    mid, end := head, head
    for mid != nil {
        end = mid.Next
        mid.Next = front
        front, mid = mid, end
    }
    return front
}
*/

/*
solution:
- recursive method:

func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	if head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
*/