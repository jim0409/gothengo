package main

import "fmt"

/*
Given an unsorted array of integers,
find the length of longest continuous increasing subsequence (subarray).

Example 1:
Input: [1,3,5,4,7]
Output: 3
Explanation:
The longest continuous increasing subsequence is [1,3,5], its length is 3.
Even though [1,3,5,7] is also an increasing subsequence,
it's not a continuous one where 5 and 7 are separated by 4.

Example 2:
Input: [2,2,2,2,2]
Output: 1
Explanation:
The longest continuous increasing subsequence is [2], its length is 1.

Note: Length of the array will not exceed 10,000.
*/

func findLengthOfLCIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	// 用來計算連續遞增數，以及暫存連續遞增數字
	count := 1
	tmpCount := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			tmpCount++
		} else {
			if tmpCount > count {
				count = tmpCount
			}
			tmpCount = 1
			continue
		}
	}

	if tmpCount > count {
		count = tmpCount
	}

	return count
}

func main() {
	input1 := []int{2, 2, 2, 2, 2}
	fmt.Println(findLengthOfLCIS(input1))
	input2 := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(input2))
	input3 := []int{1, 3, 5, 7}
	fmt.Println(findLengthOfLCIS(input3))
	input4 := []int{1, 3, 5, 4, 2, 3, 4, 5}
	fmt.Println(findLengthOfLCIS(input4))
}

/*
solution:
- https://leetcode.com/problems/longest-continuous-increasing-subsequence/discuss/394491/Go-golang-solution

func findLengthOfLCIS(nums []int) int {

    longest := 0
    res := 0

    if len(nums) == 0 { return 0 }

    for i := 1; i < len(nums); i++ {
        if nums[i-1] < nums[i] {
            longest++
            if longest > res {
                res = longest
            }
        } else {
            longest = 0
        }
    }

    return res + 1

}
*/
