package main

import "strconv"

/*
Given an integer, return its base 7 string representation.

Example 1:
Input: 100
Output: "202"

Example 2:
Input: -7
Output: "-10"

Note: The input will be in range of [-1e7, 1e7].
*/

// https://leetcode.com/problems/base-7/discuss/468757/go-clean-code-0ms-beats-100
func convertToBase7(num int) string {
	// 轉負為正
	if num < 0 {
		return "-" + convertToBase7(-num)
	}

	// 小於7不做處理
	if num < 7 {
		return strconv.Itoa(num)
	}

	// 返回該數對 7的商數+7的餘數
	return convertToBase7(num/7) + strconv.Itoa(num%7)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/base-7/discuss/464967/Go-double-100


func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	sign := ""
	if num < 0 {
		num *= -1
		sign = "-"
	}
	base7 := ""
	for num != 0 {
		base7 = strconv.Itoa(num % 7) + base7
		num /= 7
	}
	return sign+base7
}
*/
