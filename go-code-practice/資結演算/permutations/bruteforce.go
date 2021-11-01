package main

import "fmt"

func bruteForceHelper(nums []int, n int, ans *[][]int) {
	if n == 1 {
		*ans = append(*ans, copySlice(nums))
		return
	}

	for i := 0; i < n; i++ {
		nums[i], nums[n-1] = nums[n-1], nums[i]
		bruteForceHelper(nums, n-1, ans)
		nums[i], nums[n-1] = nums[n-1], nums[i]
	}
}

// BruteForce 通过暴力法生成一个序列的全部排列
func BruteForce(nums []int) [][]int {
	ans := make([][]int, 0, len(nums))

	// 將原始的 []int 陣列 以及 預計回傳的 ans 陣列 指標傳入
	bruteForceHelper(nums, len(nums), &ans)
	return ans
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(BruteForce(arr))
}

func copySlice(nums []int) []int {
	n := make([]int, len(nums), len(nums))
	copy(n, nums)
	return n
}
