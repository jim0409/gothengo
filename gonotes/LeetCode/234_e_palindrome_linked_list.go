package main

/*
Given a singly linked list, determine if it is a palindrome.

Example 1:
Input: 1->2
Output: false

Example 2:
Input: 1->2->2->1
Output: true

Follow up:
Could you do it in O(n) time and O(1) space?
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// 當串列為 nil 或一個數的時候，回傳 true
	if head == nil || head.Next == nil {
		return true
	}

	// 要找到 鏈結表 的中間點 p1
	// 利用 p2 來抓取 p1
	// 讓 p2 的 迭代速度為 p1 的兩倍；
	// head, head.Next, ... p1, p1.Next ... , p2, p2.Next
	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	var front *ListNode
	// 透過 mid節點 算出 end節點
	mid, end := p1.Next, p1.Next

	for mid != nil {
		end = mid.Next
		mid.Next = front
		front, mid = mid, end
	}

	// front 為一個從 end 開始算的 ListNode
	for front != nil {
		if head.Val != front.Val {
			return false
		}
		head = head.Next
		front = front.Next
	}
	return true
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/palindrome-linked-list/discuss/463021/Go-12ms-96

func isPalindrome(head *ListNode) bool {
	values := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Val)
	}
	for i := 0; i < len(values) / 2; i++ {
		if values[i] != values[len(values)-1-i] {
			return false
		}
	}
	return true
}
*/
