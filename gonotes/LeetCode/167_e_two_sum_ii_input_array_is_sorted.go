package main

/*
給定一個有序的數列，該數列以升冪排序。試圖找到兩個數列元素和符合目標值。
Given an array of integers that is already sorted in ascending order,
find two numbers such that they add up to a specific target number.

找到目標值後回傳兩該數列的index。
The function twoSum should return indices of the two numbers
such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
預計回傳的index1以及index2均非0
You may assume that each input would have exactly one solution and you may not use the same element twice.
假設輸入的inputs arr_num恰巧有一組解


Example:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
*/

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	if n == 0 {
		return nil
	}
	var res []int
	i := 0
	j := n - 1

	// 從最小及最大數列來算和，和小於目標值"左邊"index右移，反之"左邊"index左移
	// binary search: 如果總和大於目標值 j-- / 反之 i++
	for i < j && j > 0 {
		if numbers[i]+numbers[j] == target {
			res = append(res, i+1)
			res = append(res, j+1)
			return res
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return res
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/discuss/344776/Go-Two-pointer-solution


func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for {
		if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
}
*/
