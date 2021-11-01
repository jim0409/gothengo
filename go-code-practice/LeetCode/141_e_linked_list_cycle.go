package main

/*
Given a linked list, determine if it has a cycle in it.
To represent a cycle in the given linked list,
we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to.
If pos is -1, then there is no cycle in the linked list.


Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the second node.


Example 2:

Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the first node.


Example 3:

Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.


Follow up:
Can you solve it using O(1) (i.e. constant) memory?
*/

func hasCycle(head *ListNode) bool {
	// 製作一個 dict 確認是否存在該node
	dict := make(map[*ListNode]struct{})

	// 當 head 不為 nil 時，依序將 head 加入 dict 中
	for head != nil {

		// 只要有辦法從 dict 中取到 dict[head] 的值，就可以返回 true
		if _, ok := dict[head]; ok {
			return true
		}

		dict[head] = struct{}{}
		head = head.Next
	}
	return false
}

func main() {

}

/*
solution:
https://leetcode.com/problems/linked-list-cycle/discuss/526451/Go

-- two pointers

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    p1, p2 := head, head.Next

    // https://zh.wikipedia.org/wiki/Floyd-Warshall%E7%AE%97%E6%B3%95
    for p1 != p2 {
        if p2 == nil || p2.Next == nil {
            return false
        }
        p1 = p1.Next
        p2 = p2.Next.Next
    }
    return true
}


-- hash

func hasCycle(head *ListNode) bool {
    dict := make(map[*ListNode]struct{})
    for head != nil {
        if _, ok := dict[head]; ok {
            return true
        }
        dict[head] = struct{}{}
        head = head.Next
    }
    return false
}
*/
