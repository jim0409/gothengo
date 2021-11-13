package main

import (
	"fmt"
	"math"
)

/*
Given an integer array,
you need to find one continuous subarray that
if you only sort this subarray in ascending order,
then the whole array will be sorted in ascending order, too.

You need to find the shortest such subarray and output its length.

Example 1:
Input: [2, 6, 4, 8, 10, 9, 15]
Output: 5
Explanation:
You need to sort [6, 4, 8, 10, 9] in ascending order
to make the whole array sorted in ascending order.

Note:
Then length of the input array is in range [1, 10,000].
The input array may contain duplicates, so ascending order here means <=.
*/

// 找到搜尋的頭。從陣列的第一個數字開始向後比，找到非遞增的數列。就往前退一格
// 從尾部開始找，從陣列的最後一個數字開始向前比，找到第一個非遞減的數列+1
// 計算兩個數字的差
// https://leetcode.com/problems/shortest-unsorted-continuous-subarray/discuss/485597/go-clean-code-16ms-beats-100
func findUnsortedSubarray(nums []int) int {
	start, end, min, max := 0, -1, math.MaxInt32, math.MinInt32
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		// 從頭部開始搜尋，找到第一個非遞減的數字記為`end`
		if nums[i] >= max {
			max = nums[i]
		} else {
			end = i
		}

		// 從尾部開始搜尋，找到第一個非遞增的數字記為`start`
		if nums[j] <= min {
			min = nums[j]
		} else {
			start = j
		}
	}
	return end - start + 1
}

func main() {
	input := []int{2, 6, 4, 8, 10, 9, 7}
	fmt.Println(findUnsortedSubarray(input))

	input2 := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println(findUnsortedSubarray(input2))
}

/*
solution:
- https://leetcode.com/problems/shortest-unsorted-continuous-subarray/discuss/661286/Go-2-passes-solution


func satisfy(a, b, c, d int) bool {
    return a <= b && b <= c && c <= d
}

func findUnsortedSubarray(nums []int) int {
    if len(nums) <= 1 {
        return 0
    }

	// step 1: find a pair of (i, j) to make sure 2 sides are ascending
    i, j := 0, len(nums) - 1
    for i < j && nums[i] <= nums[i + 1] {
        i ++
    }

    for i < j && nums[j - 1] <= nums[j] {
        j --
    }

	// prepare max and min
    min, max := nums[i], nums[j]
    if i < j {
        for p := i; p <= j; p ++ {
            if nums[p] < min {
                min = nums[p]
            }

            if nums[p] > max {
                max = nums[p]
            }
        }
    }

	// step 2: adjust (i, j) pair
    for i >= 0 && j < len(nums) && !satisfy(nums[i], min, max, nums[j]) {
        for i >= 0 && nums[i] > min {
            if nums[i] > max {
                max = nums[i]
            }

            i --
        }

        for j < len(nums) && nums[j] < max {
            if nums[j] < min {
                min = nums[j]
            }

            j ++
        }
    }

    if i == j {
        return 0
    }

    return j - i - 1
}

*/
