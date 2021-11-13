package main

import "fmt"

/*
Given an array nums with n integers,
your task is to check if it could become non-decreasing by modifying at most 1 element.

We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based)
such that (0 <= i <= n - 2).


Example 1:
Input: nums = [4,2,3]
Output: true
Explanation:
You could modify the first 4 to 1 to get a non-decreasing array.


Example 2:
Input: nums = [4,2,1]
Output: false
Explanation:
You can't get a non-decreasing array by modify at most one element.
*/

// func checkPossibility(nums []int) bool {
// 	var res int
// 	l := len(nums)
// 	for i := 0; i < l-1; i++ {
// 		if nums[i] > nums[i+1] {
// 			res++
// 		}
// 	}
// 	return res <= 1
// }

func checkPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	modified := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			continue
		}
		if modified {
			return false
		}
		if i > 0 && nums[i-1] > nums[i+1] {
			nums[i+1] = nums[i]
		}
		modified = true
	}
	return true
}

func main() {
	input1 := []int{3, 4, 2, 3}
	fmt.Println(checkPossibility(input1))
}

/*
solution:
- https://leetcode.com/problems/non-decreasing-array/discuss/359088/Go-solution-87-100

func checkPossibility(nums []int) bool {
	breakpoint := false // mark has break or not
	for i, v := range nums {
		if i == 0 || nums[i-1] <= v {
			continue
		}

		if breakpoint {   // found second break, return false
			return false
		}

		breakpoint = true  // found break, mark
		if i == 1 || i == len(nums)-1 {  // 1 and len(nums)-1 is OK
			continue
		} else if nums[i-2] <= v || nums[i-1] <= nums[i+1] {
			continue
		} else {
			return false
		}
	}
	return true
}
*/

/*
solution:
- https://leetcode.com/problems/non-decreasing-array/discuss/492427/go-clean-code-16ms-beats-100

func checkPossibility(nums []int) bool {
	nums = append([]int{math.MinInt64}, nums...)
	nums = append(nums, math.MaxInt64)
	for flag, i := false, 1; i < len(nums)-2; i++ {
		if nums[i] <= nums[i+1] {
			continue
		}
		if flag || (nums[i] > nums[i+2] && nums[i-1] > nums[i+1]) {
			return false
		}
		flag = true
	}
	return true
}
*/
