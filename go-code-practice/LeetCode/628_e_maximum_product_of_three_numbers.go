package main

import "math"

/*
Given an integer array, find three numbers
// 可能為 +/ -/ -
whose product is maximum and output the maximum product.

Example 1:
Input: [1,2,3]
Output: 6


Example 2:
Input: [1,2,3,4]
Output: 24


Note:
1. The length of the given array will be in range [3,104]
and all elements are in the range [-1000, 1000].

2. Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.
*/

// https://leetcode.com/problems/maximum-product-of-three-numbers/discuss/486465/go-clean-code-28ms-beats-100
func maximumProduct(nums []int) int {
	max1, max2, max3, min1, min2 := math.MinInt32, math.MinInt32, math.MinInt32, math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num > max1 {
			max1, max2, max3 = num, max1, max2
		} else if num > max2 {
			max2, max3 = num, max2
		} else if num > max3 {
			max3 = num
		}
		if num < min1 {
			min1, min2 = num, min1
		} else if num < min2 {
			min2 = num
		}
	}
	return int(math.Max(float64(max1*max2*max3), float64(max1*min1*min2)))
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/maximum-product-of-three-numbers/discuss/389312/Go-golang-sorted-and-not-sorted-solutions

[sort-slice]
func maximumProduct(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	a := nums[0] * nums[1] * nums[2]
	b := nums[0] * nums[len(nums)-1] * nums[len(nums)-2]
	if a > b { return a }
	return b
}


[find maxs3 or max1 with mins2]
func maximumProduct(nums []int) int {

	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
    min1, min2 := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num >= max1 {
			max3, max2, max1 = max2, max1, num
		} else if num >= max2 {
			max3, max2 = max2, num
		} else if num > max3 {
			max3 = num
		}
		if num <= min1 {
			min2, min1 = min1, num
		} else if num < min2 {
			min2 = num
		}
	}
    if max1 * max2 * max3 > min1 * min2 * max1 { return max1 * max2 * max3 }
	return min1 * min2 * max1
}
*/
