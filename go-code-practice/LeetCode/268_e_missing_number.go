package main

import "fmt"

/*
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n,
find the one that is missing from the array.

Example 1:
Input: [3,0,1]
Output: 2

Example 2:
Input: [9,6,4,2,3,5,7,0,1]
Output: 8

Note:
Your algorithm should run in linear runtime complexity.
Could you implement it using only constant extra space complexity?
*/

// wrong ...
/*
func missingNumber(nums []int) int {
	n := len(nums)
	m := make(map[int]bool, n)

	for i := 0; i < n; i++ {
		m[nums[i]] = true
	}

	for i, j := range m {
		if j {
			return i
		}
	}

	return 0
}
*/

/*
	- 比較總和
	由題目得知，長度為 n 的數列中，會恰巧缺少數列中的其中一個值
	例如長度 5 ，只會從 [0,1,2,3,4] 中選擇一個來缺少
	因此，缺少值的陣列總和，與實際 0+...+5的值，會恰好差"缺少值"
*/
func missingNumber(nums []int) int {

	currentSum := 0
	for i := 0; i < len(nums); i++ {
		currentSum = currentSum + nums[i] - i
	}
	return -(currentSum - len(nums))
}

func main() {
	input := []int{3, 0, 1}
	fmt.Println(missingNumber(input))
}

/*
solution:
- https://leetcode.com/problems/missing-number/discuss/258669/100-use-go


func missingNumber(nums []int) int {

    currentSum := 0
    for i:=0;i<len(nums);i++{
        currentSum = currentSum + nums[i] -i
    }
    return -(currentSum - len(nums))
}
*/
