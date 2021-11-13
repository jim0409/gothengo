package main

import "fmt"

/*
Given an array, rotate the array to the right by k steps, where k is non-negative.


Example 1:
Input: [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]


Example 2:
Input: [-1,-100,3,99] and k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]

Note:
Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?
*/

// func rotate(nums []int, k int) {
// 	lenNums := len(nums)
// 	nums1 := nums[:lenNums-k]
// 	nums2 := nums[lenNums-k:]
// 	fmt.Println(nums1, nums2)
// 	for i := 0; i < k; i++ {
// 		nums[i] = nums2[i]
// 	}

// 	for i := 0; i < lenNums-k; i++ {
// 		fmt.Println(nums1[i])
// 	}
// 	/*
// 		需要用make額外宣告，因為 array[:index] 是回傳 pointer
// 		[1 2 3 4] [5 6 7]
// 		5
// 		6
// 		7
// 		4
// 	*/
// }

func rotate(nums []int, k int) {

	n := len(nums)
	rot_arr := make([]int, n)
	for i := 0; i < n; i++ {
		rot_arr[(i+k)%n] = nums[i]
	}
	copy(nums, rot_arr)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

/*
solution:
- https://leetcode.com/problems/rotate-array/discuss/211401/Go-solution


func rotate(nums []int, k int)  {

    n := len(nums)
    rot_arr := make([]int, n)
    for i:= 0; i<n; i++ {
        rot_arr[(i+k)%n] = nums[i]
    }
    copy(nums,rot_arr)
}
*/
