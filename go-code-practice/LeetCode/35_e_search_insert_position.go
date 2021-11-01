package main

import "fmt"

/*
Given a sorted array and a target value, return the index if the target is found.

If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0

*/

func searchInsert(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	if target < nums[0] {
		return 0
	}

	left := 0
	// right := len(nums) - 1
	for i, j := range nums {
		if j == target {
			return i
		}

		if target > j {
			left = i + 1
			if nums[i+1] > target {
				return left
			}
		}

		// if j < target {
		// right = i
		// }
	}

	return 0
}

func main() {
	arr1 := []int{1, 3, 5, 6}
	t1 := 5

	fmt.Println(searchInsert(arr1, t1))

	// arr2 := []int{1, 3, 5, 6}
	// t2 := 5

	// arr3 := []int{1, 3, 5, 6}
	// t3 := 7

	// arr4 := []int{1, 3, 5, 6}
	// t4 := 0

	arr5 := []int{1, 3}
	t5 := 2
	fmt.Println(searchInsert(arr5, t5))

	arr6 := []int{1, 3, 5, 6}
	t6 := 2
	fmt.Println(searchInsert(arr6, t6))

}

/*
參考解答
https://leetcode.com/problems/search-insert-position/discuss/315526/Go-4ms

func searchInsert(nums []int, target int) int {
    l := 0
    r := len(nums) - 1
    idx := 0
    if target > nums[r]{
        return len(nums)
    }else if target < nums[l]{
        return 0
    }else{
        for l < r{
            idx = (l+r)/2
            if nums[idx] == target{
                return idx
            }else if target > nums[idx]{
                l++
            }else{
                r--
            }
        }
    }
    return l
}

*/
