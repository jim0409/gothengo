package main

import "fmt"

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		// 如果要進行排列組合的長度為1 e.g. []int{1}，那麼不需要做置換，直接將copy出來的值返回給res
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			// 從 n = len(arr) 開始，i = 0, 1 ...
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				// 判斷要置換的長度是否為偶數or基數
				if n%2 == 1 { // n 為 基數
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(permutations(arr))
}
