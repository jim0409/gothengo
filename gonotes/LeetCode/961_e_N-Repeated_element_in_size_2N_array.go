package main

/*
In a array nums of size 2 * n, there are n + 1 unique elements,
and exactly one of these elements is repeated n times.

Return the element repeated n times.


Example 1:
Input: nums[1,2,3,3]
Output: 3

Example 2:
Input: nums[2,1,2,5,3,2]
Output: 2

Example 3:
Input: nums[5,1,5,2,5,3,5,4]
Output: 5


Note:
4 <= nums.length <= 10000
0 <= nums[i] < 10000
nums.length is even
*/

// func repeatedNTimes(nums []int) int {
// 	var res int
// 	var m = make(map[int]int, len(nums))
// 	for _, j := range nums {
// 		if v, ok := m[j]; ok {
// 			m[j] = v + 1
// 		} else {
// 			m[j] = 1
// 		}
// 	}
//
// 	for i, j := range m {
// 		if j != 1 {
// 			res = i
// 			break
// 		}
// 	}

// 	return res
// }

// - https://leetcode.com/problems/n-repeated-element-in-size-2n-array/discuss/220994/100-use-go
func repeatedNTimes(A []int) int {
	var mp [10000]int
	l := len(A)
	for i := 0; i < l; i++ {
		mp[A[i]]++
		if mp[A[i]] > 1 {
			return A[i]
		}
	}
	return 0
}

func main() {}

/*
refer:
- https://leetcode.com/problems/n-repeated-element-in-size-2n-array/discuss/581988/go-clean-code-20ms-beats-100

func repeatedNTimes(A []int) int {
	for i := 2; i < len(A); i++ {
		if A[i] == A[i-1] || A[i] == A[i-2] {
			return A[i]
		}
	}
	return A[0]
}
*/
