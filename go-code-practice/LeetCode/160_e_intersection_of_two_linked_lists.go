package main

/*
Write a program to find the node at which the intersection of two singly linked lists begins.

For example, the following two linked lists:


Example 1:
Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
Output: Reference of the node with value = 8

Input Explanation: The intersected node's value is 8
(note that this must not be 0 if the two lists intersect).

From the head of A, it reads as [4,1,8,4,5].
From the head of B, it reads as [5,0,1,8,4,5].

There are 2 nodes before the intersected node in A;
There are 3 nodes before the intersected node in B.


Example 2:
Input: intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Reference of the node with value = 2

Input Explanation: The intersected node's value is 2
(note that this must not be 0 if the two lists intersect).

From the head of A, it reads as [0,9,1,2,4].
From the head of B, it reads as [3,2,4].

There are 3 nodes before the intersected node in A;
There are 1 node before the intersected node in B.


Example 3:
Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: null

Input Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5].
Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.


Notes:

If the two linked lists have no intersection at all, return null.
The linked lists must retain their original structure after the function returns.
You may assume there are no cycles anywhere in the entire linked structure.
Your code should preferably run in O(n) time and use only O(1) memory.
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 定義出 nodeA 以及 nodeB 方便後面計算出 lengthA & lengthB
	nodeA, nodeB := headA, headB
	lengthA, lengthB := 0, 0
	for nodeA != nil {
		nodeA = nodeA.Next
		lengthA++
	}
	for nodeB != nil {
		nodeB = nodeB.Next
		lengthB++
	}

	// 宣告 longer & shorter - 以及 diff
	longer, shorter, diff := headA, headB, lengthA-lengthB
	if lengthA < lengthB {
		longer, shorter = headB, headA
		diff = lengthB - lengthA
	}
	// => 令兩個 linked list 長度一致，讓後面的for可以去同時比較?(不考慮 A=[1,2,1] && B=[3,1] 交集在 A[0] 的情況)
	for i := 0; i < diff; i++ {
		longer = longer.Next
	}

	for {
		// 當 longer & shorter 的 ListNode 一致(可能同值或者null) 就返回
		if longer == shorter {
			return longer
		}
		longer = longer.Next
		shorter = shorter.Next
	}
}
func main() {

}

/*
solution:
- https://leetcode.com/problems/intersection-of-two-linked-lists/discuss/351825/Go-O(n%2Bm)-2-different-solutions


// With calculation diff between lengths of lists

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	lenA, lenB := lengthOfList(headA), lengthOfList(headB)
	if lenA < lenB {
		headA, headB = headB, headA
		lenA, lenB = lenB, lenA
	}
	for i := 0; i < lenA-lenB; i++ {
		headA = headA.Next
	}

	for headA != nil && headB != nil && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}

func lengthOfList(l *ListNode) int {
	var length int
	for l != nil {
		l = l.Next
		length++
	}
	return length
}



// Without it

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p, q := headA, headB
	for p != nil || q != nil {
		if p == q {
			return p
		}

		if p == nil {
			p = headB
			q = q.Next
			continue
		}

		if q == nil {
			q = headA
			p = p.Next
			continue
		}

		p = p.Next
		q = q.Next
	}
	return nil
}

*/
