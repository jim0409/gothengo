package main

/*
Given a non-empty, singly linked list with head node head, return a middle node of linked list.

If there are two middle nodes, return the second middle node.


Example 1:
Input: [1,2,3,4,5]
Output: Node 3 from this list (Serialization: [3,4,5])
The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).
Note that we returned a ListNode object ans, such that:
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.

Example 2:
Input: [1,2,3,4,5,6]
Output: Node 4 from this list (Serialization: [4,5,6])
Since the list has two middle nodes with values 3 and 4, we return the second one.


Note:
The number of nodes in the given list will be between 1 and 100.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 求中間節點值，如果有兩個中間節點則提取第二個節點值當作中間節點值
func middleNode(head *ListNode) *ListNode {
	// 宣告兩個指標，讓第二個指標的移動速度是第一個指標移動速度的兩倍，那麼slow就是中間節點值
	slow, fast := head, head
	// 只要第二個指標的下一個，及下下一個節點值不為nil則持續迭代
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 如過快指標的下一個，及下下一個節點值均為nil則代表有兩個中間節點值，回傳第二個
	if fast.Next != nil && fast.Next.Next == nil {
		return slow.Next
	}

	return slow
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/middle-of-the-linked-list/discuss/348588/Go-Two-pointer

func middleNode(head *ListNode) *ListNode {
  dummy := new(ListNode)
  dummy.Next = head
  slow, fast := dummy, dummy.Next

  for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
  }

  return slow.Next
}
*/
