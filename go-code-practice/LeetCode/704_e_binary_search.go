package main

import "fmt"

/*
Given a sorted (in ascending order) integer array nums of n elements and a target value,
write a function to search target in nums.

If target exists, then return its index, otherwise return -1.


Example 1:
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation:
9 exists in nums and its index is 4

Example 2:
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation:
2 does not exist in nums so return -1


Note:
1. You may assume that all elements in nums are unique.
2. n will be in the range [1, 10000].
3. The value of each element in nums will be in the range [-9999, 9999].
*/

// func search(nums []int, target int) int {
// 	for i, j := range nums {
// 		if j == target {
// 			return i
// 		}
// 	}

// 	return -1
// }

func search(nums []int, target int) int {
	for i, j := 0, len(nums)-1; i <= j; {
		m := i + (j-i)/2
		if nums[m] < target {
			i = m + 1
		} else if nums[m] > target {
			j = m - 1
		} else {
			return m
		}
	}
	return -1
}

func main() {
	input1 := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(input1, 9))
}

/*
solution:
- https://leetcode.com/problems/binary-search/discuss/702059/Commented-solution-in-Go

func search(nums []int, target int) int {
    // It's important to be clear whether the high bound is inclusive
	// or exclusive. In this case it's included and we're searching
	// in the range [low, high].
    low := 0
    high := len(nums) - 1

	// Pay special attention to the condition, in combination with
	// how the bounds are updated in the loop. If you're not careful,
	// it's possible to create an infinite loop, especially in the case
	// when the target is not found and low and high cannot get
	// any closer to each other.
	//
	// Also consider what will happen when the input has
	// zero elements or one element.
    for low <= high {
	    // In Go, with ints, int division is used.
        mid := (low + high) / 2
        if nums[mid] == target {
		    // Found it!
            return mid
        } else if nums[mid] < target {
		    // mid is eliminated, the lowest possible index
			// is higher than low.
            low = mid + 1
        } else if nums[mid] > target {
		    // high is eliminated, the highest possible index
			// is lower than high.
            high = mid - 1
        }
    }
    return -1
}
*/
