package main

/*
Remove all elements from a linked list of integers that have value val.

Example:

Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// 製作一個dummy，避免改到 head 的值 {Val: 0, Next: head}
	// dummy 比起 head，前頭多一個節點
	dummy := &ListNode{Next: head}
	// 宣告 i,j 來分別操作 dummy & head
	i, j := dummy, head

	// 如果 head 非空就往下遞迴求解
	for j != nil {
		// 當 head.Val = val 時
		if j.Val == val {
			// dummy.Next = head.Next <=> head = head.Next
			// 註: 不能直接用 head, 因為會改動到下面的賦值
			i.Next = j.Next
			// head = head
			j = i.Next
		} else {
			i = i.Next
			j = j.Next
		}
	}
	return dummy.Next
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/remove-linked-list-elements/discuss/142417/Iterative-Go-solution


func removeElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{-1, head}
    p := dummy
    for p.Next != nil {
        if p.Next.Val == val {
            p.Next = p.Next.Next
            continue
        }
        p = p.Next
    }
    return dummy.Next
}
*/
