package main

import "fmt"

/*
In a given integer array nums, there is always exactly one largest element.
Find whether the largest element in the array is at least twice as much as every other number in the array.
判斷陣列中是否存在該元素(比其他元素的2倍還大，有則返回該元素位置。否則返回-1)
If it is, return the index of the largest element, otherwise return -1.

Example 1:
Input: nums = [3, 6, 1, 0]
Output: 1
Explanation:
6 is the largest integer, and for every other number in the array x,
6 is more than twice as big as x.  The index of value 6 is 1, so we return 1.


Example 2:
Input: nums = [1, 2, 3, 4]
Output: -1
Explanation:
4 isn't at least as big as twice the value of 3, so we return -1.


Note:
nums will have a length in the range [1, 50].
Every nums[i] will be an integer in the range [0, 99].
*/

// func dominantIndex(nums []int) int {
// 	if len(nums) <= 1 {
// 		return len(nums)
// 	}

// 	tmpIndex := 0

// 	for i, j := range nums {
// 		if j >= nums[tmpIndex] {
// 			tmpIndex = i
// 		}
// 	}

// 	return tmpIndex
// }

func dominantIndex(nums []int) int {
	highest, secHighest := 0, 0
	key := 0
	for k, v := range nums {
		if v > highest {
			secHighest = highest
			highest = v
			key = k
		} else if v > secHighest {
			secHighest = v
		}
	}

	if 2*secHighest <= highest {
		return key
	} else {
		return -1
	}
}

func main() {
	input1 := []int{3, 6, 1, 0}
	fmt.Println(dominantIndex(input1))
}

/*
solution:
- https://leetcode.com/problems/largest-number-at-least-twice-of-others/discuss/351171/Go-0ms-(faster-than-100)


func dominantIndex(nums []int) int {
  max, second, imax := nums[0], 0, 0
  for i := 1; i < len(nums); i++ {
    if nums[i] > max {
      second = max
      max, imax = nums[i], i
    } else if nums[i] > second {
      second = nums[i]
    }
  }
  if max >= second * 2 {
    return imax
  }
  return -1
}
*/
