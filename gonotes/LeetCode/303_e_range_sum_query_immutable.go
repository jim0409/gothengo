package main

import "fmt"

/*
Given an integer array nums,
find the sum of the elements between indices i and j (i ≤ j), inclusive.

Example:

Given nums = [-2, 0, 3, -5, 2, -1]
sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3

Note:
You may assume that the array does not change.
There are many calls to sumRange function.
*/

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{
			arr: []int{},
		}
	}
	tmp := make([]int, len(nums))
	tmp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		tmp[i] = tmp[i-1] + nums[i]
	}

	return NumArray{
		arr: tmp,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.arr[j]
	}
	return this.arr[j] - this.arr[i-1]
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	fmt.Println(obj.SumRange(0, 2))
}

/*
solution:
- https://leetcode.com/problems/range-sum-query-immutable/discuss/458583/go-16-line-code-32ms-beats-98


type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{
		nums: append([]int{0}, nums...),
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.nums[j+1] - this.nums[i]
}
*/
