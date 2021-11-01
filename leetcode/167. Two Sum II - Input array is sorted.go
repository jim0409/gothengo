package main

import (
	"fmt"
	"strings"
	"strconv"
	"math"
)

/***************************************/
//some common data struct and function using in leetcode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

type Point struct {
	X int
	Y int
}

func useLib(){
	fmt.Println(strconv.Itoa(1))
	fmt.Println(strings.Compare("1","2"))
	fmt.Println(math.Abs(1.0))
}

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = nums[0]
	ch := make(chan *TreeNode, len(nums))
	ch <- root
	nums = nums[1:]
	for i := 0; i < len(nums); i++ {
		tree := <-ch
		if nums[i] == -1 {
			tree.Left = nil
		} else {
			tree.Left = &TreeNode{
				Val: nums[i],
			}
			ch <- tree.Left
		}
		i++
		if i == len(nums) || nums[i] == -1 {
			tree.Right = nil
		} else {
			tree.Right = &TreeNode{
				Val: nums[i],
			}
			ch <- tree.Right
		}
	}
	return root
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	root := &ListNode{
		Val: nums[0],
	}
	tmp := root
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{
			Val: nums[i],
		}
		tmp = tmp.Next
	}
	return root
}
/**************************************************************/
func twoSum(nums []int, target int) []int {
    for i:=0;i<len(nums);i++{
		t := target - nums[i]
		for left,right:= i+1,len(nums)-1;left<=right;{
			mid := (left + right)/2
			if t == nums[mid]{
				return []int{i,mid}
			}
			if nums[mid] <t{
				left = mid +1
			}else{
				right = mid -1
			}
		}
	}
	return []int{}
}
func main() {
	fmt.Println(twoSum([]int{2,7,11,15},9))
}
