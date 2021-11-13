package main

import (
	"fmt"
)

/*
  Definition for singly-linked list.
  type ListNode struct {
      Val int
      Next *ListNode
  }
  given two non-empty linked lists
  representing tow non-negative integers.

  The digits are stored in reverse order and each of
  their nodes contain a single digit.

  Add the tow numbers and return it as a linked list
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1. 判斷 l1 與 l2 是否為空，否則繼續迭代
2. 判斷 l3 的單一節點是否大於 10，是則進位
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{Val: l1.Val + l2.Val, Next: nil}
	curNode := l3

	// 第一個 for 把l1 與 l2 相加起來
	for l1.Next != nil || l2.Next != nil {
		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1.Val = 0 // 如果 link已經到底了，爾後就算0
		}

		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2.Val = 0
		}

		tmpNode := &ListNode{Val: l1.Val + l2.Val, Next: nil}
		curNode.Next = tmpNode
		curNode = curNode.Next
	}

	curNode = l3
	if curNode.Val >= 10 {
		curNode.Val = curNode.Val % 10
		if curNode.Next == nil {
			curNode.Next = &ListNode{Val: 1, Next: nil}
		} else {
			curNode.Next.Val++
		}
	}

	for curNode.Next != nil {
		if curNode.Next.Val >= 10 {
			curNode.Next.Val = curNode.Next.Val % 10
			if curNode.Next.Next == nil {
				curNode.Next.Next = &ListNode{Val: 1, Next: nil}
			} else {
				curNode.Next.Next.Val++
			}
		}
		curNode = curNode.Next
	}

	return l3
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

func main() {
	// given Input : (2 -> 4 -> 13) + (5 -> 6 -> 4)
	// output : 7 -> 0 -> 8
	// Explanation 342 + 465 = 807
	a1 := []int{2, 4, 3}
	a2 := []int{5, 6, 4}
	l1 := createLists(a1)
	l2 := createLists(a2)
	resList := addTwoNumbers(l1, l2)
	dumpList(resList)
	fmt.Println()

	a3 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	a4 := []int{5, 6, 4}
	l3 := createLists(a3)
	l4 := createLists(a4)
	resList2 := addTwoNumbers(l3, l4)
	dumpList(resList2)
	fmt.Println()

	a5 := []int{5}
	a6 := []int{5}
	l5 := createLists(a5)
	l6 := createLists(a6)
	resList3 := addTwoNumbers(l5, l6)
	dumpList(resList3)
	fmt.Println()

	a7 := []int{1, 8}
	a8 := []int{0}
	l7 := createLists(a7)
	l8 := createLists(a8)
	resList4 := addTwoNumbers(l7, l8)
	dumpList(resList4)
	fmt.Println()
}

func dumpList(l *ListNode) {
	for l != nil {
		fmt.Printf("%v ", l.Val)
		l = l.Next
	}
}

/*
	錯誤，會造成溢出
*/
/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// resultListNode := &ListNode{}

	// firstN, secN := l1.Val, l2.Val
	firstN := float64(l1.Val)
	secN := float64(l2.Val)
	lCount, rCount := 10, 10

	// 這樣的設計有可能造成golang的int溢位
	for l1.Next != nil {
		firstN += float64(l1.Next.Val * lCount)
		l1 = l1.Next
		lCount = 10 * lCount
	}

	for l2.Next != nil {
		secN += float64(l2.Next.Val * rCount)
		l2 = l2.Next
		rCount = 10 * rCount
	}

	sumN := fmt.Sprintf("%v", firstN+secN)

	resultArr := []int{}
	for i := len(sumN); i > 0; i-- {
		val, _ := strconv.Atoi(sumN[i-1 : i])
		resultArr = append(resultArr, val)
	}

	return createLists(resultArr)
}
*/
