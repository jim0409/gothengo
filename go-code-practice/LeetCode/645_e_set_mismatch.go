package main

import "fmt"

/*
The set `S` originally contains numbers from 1 to n.

But unfortunately, due to the data error,
`one` of the numbers in the set got duplicated to another number in the set,
which results in repetition of one number and loss of another number.

Given an array nums representing the data status of this set after the error.
Your task is to firstly find the number occurs twice and then find the number that is missing.
Return them in the form of an array.


Example 1:
Input: nums = [1,2,2,4]
Output: [2,3]


Note:
The given array size will in the range [2, 10000].
The given array's numbers won't have any order.
*/

// func findErrorNums(nums []int) []int {
// 	var pos int
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] == nums[i-1] {
// 			pos = i
// 		}
// 	}

// 	return []int{pos, pos + 1}
// }

// - https://leetcode.com/problems/set-mismatch/discuss/556014/GO-Using-Constant-Space
func findErrorNums(nums []int) []int {
	dup := 0
	for _, v := range nums {
		n := abs(v)
		if nums[n-1] < 0 {
			dup = n
		} else {
			nums[n-1] *= -1
		}

	}

	mis := 0
	for i, v := range nums {
		if v > 0 {
			mis = i + 1
			break
		}
	}

	return []int{dup, mis}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	input := []int{1, 2, 2, 4}
	fmt.Println(findErrorNums(input))
}

/*
solution:
- https://leetcode.com/problems/set-mismatch/discuss/486610/go-clean-code-20ms-beats-100

func findErrorNums(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for ; nums[i] != nums[nums[i]-1]; nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i] {}
	}
	for i, num := range nums {
		if num != i+1 {
			return []int{num, i + 1}
		}
	}
	return nil
}
*/
