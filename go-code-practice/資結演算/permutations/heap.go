package main

import "fmt"

func Heap(arr []int) [][]int {
	res := [][]int{}
	helper(arr, len(arr), &res)
	return res
}

func helper(arr []int, n int, res *[][]int) {
	// 如果要進行排列組合的長度為1 e.g. []int{1}，那麼不需要做置換，直接將copy出來的值返回給res
	if n == 1 {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*res = append(*res, tmp)
	} else {
		// 從 n = len(arr) 開始，i = 0, 1 ...
		for i := 0; i < n; i++ {
			// 判斷要置換的長度是否為偶數or基數
			helper(arr, n-1, res)
			if n%2 == 1 { // 如果 n 是基數
				arr[i], arr[n-1] = arr[n-1], arr[i]

			} else {
				arr[0], arr[n-1] = arr[n-1], arr[0]

			}
		}
	}
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(Heap(arr))
}
