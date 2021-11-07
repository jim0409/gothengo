package main

/*
Given an integer array nums sorted in non-decreasing order,

return an array of the squares of each number sorted in non-decreasing order.

Example 1:
Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Example 2:
Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]


Constraints:
1 <= nums.length <= 10^4
-10^4 <= nums[i] <= 10^4
nums is sorted in non-decreasing order.
*/

/*
// 第一次的寫法 n^2
func squares(i int) int {
	return i * i
}

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = squares(nums[i])
	}
	sort.Ints(nums)
	return nums
}
*/

// O(n)
// refer: https://leetcode.com/problems/squares-of-a-sorted-array/discuss/309345/golang-code.-Beautiful-and-clear
func sortedSquares(A []int) []int {
	length := len(A)
	ans := make([]int, length)
	// 使用 2 pointer 方法
	l, r := 0, length-1
	// 用 右址 初始 index
	index := r

	// 當 右址 小於 左址 時跳出判斷
	for l <= r {

		// 如果 右值 + 左值 < 0 ，
		// ans[index] 為 A[l]*A[l] 左址++，反之
		// ans[index] 為 A[r]*A[r] 右址--
		if A[l]+A[r] < 0 {
			ans[index] = A[l] * A[l]
			l++
		} else {
			ans[index] = A[r] * A[r]
			r--
		}

		// 每做一次 index --
		index--
	}

	return ans
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/squares-of-a-sorted-array/discuss/1521532/Golang-2-pointer-O(N)

func sortedSquares(nums []int) []int {
	out := make([]int, len(nums))
	left := 0
	right := len(nums) - 1
	index := right

	for left <= right {
		lval := nums[left] * nums[left]
		rval := nums[right] * nums[right]
		if lval >= rval {
			out[index] = lval
			left++
		} else {
			out[index] = rval
			right--
		}

		index--
	}

	return out
}
*/
