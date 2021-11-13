package main

import "fmt"

/*
You are given two arrays (without duplicates) nums1 and nums2 where nums1’s elements are subset of nums2.
給定兩個 不重複整數陣列，其中 num1 是 num2 的子集合
Find all the next greater numbers for nums1's elements in the corresponding places of nums2.
找出所有 在 nums1 下的元素，該元素所對應到的 nums2 元素中，是否可以找到比該數還要大的數。可以就返回該數，否則返回-1

The Next Greater Number of a number x in nums1 is the first greater number to its right in nums2.
If it does not exist, output -1 for this number.

Example 1:
Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
Output: [-1,3,-1]
Explanation:
    For number 4 in the first array, you cannot find the next greater number for it in the second array, so output -1.  // 4對應到[...4.] 回傳3; 但是nums[3]=2沒有比4大，回傳-1
    For number 1 in the first array, the next greater number for it in the second array is 3. // 1對應到[.3...]; 可以找到nums2[1]=3，發現nums[2]=4比他大。所以回傳3
	For number 2 in the first array, there is no next greater number for it in the second array, so output -1. // 2對應到[...2]; 無法找到nums[3]之後的數字。所以直接回傳找不到-1

Example 2:
Input: nums1 = [2,4], nums2 = [1,2,3,4].
Output: [3,-1]
Explanation:
    For number 2 in the first array, the next greater number for it in the second array is 3.
    For number 4 in the first array, there is no next greater number for it in the second array, so output -1.

Note:
All elements in nums1 and nums2 are unique.
The length of both nums1 and nums2 would not exceed 1000.
*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	out := make([]int, 0)
	// for _, item := range nums1 {
	for i := 0; i < len(nums1); i++ {
		item := nums1[i]
		l := len(out)
		find := false
		for j := 0; j < len(nums2); j++ {
			// 判斷 item 是否存在於 nums2 裡面
			if nums2[j] == item {
				find = true
			}

			// 如果存在，則判斷是否有next Greater
			if find && nums2[j] > item {
				out = append(out, nums2[j])
				break
			}
		}

		if l == len(out) {
			out = append(out, -1)
		}
	}
	return out
}

func main() {
	nums11 := []int{4, 1, 2}
	nums21 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums11, nums21))
}

/*
solution:
- https://leetcode.com/problems/next-greater-element-i/discuss/467979/go-clean-code-0ms-beats-100


func nextGreaterElement(nums1 []int, nums2 []int) []int {
	data := make(map[int]int, len(nums2))
	var stack []int
	for _, num := range nums2 {
		for len(stack) != 0 && stack[len(stack)-1] < num {
			data[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	ret := make([]int, len(nums1))
	for i, num := range nums1 {
		ret[i] = -1
		if v, ok := data[num]; ok {
			ret[i] = v
		}
	}
	return ret
}
*/
