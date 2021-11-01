package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := copySlice(a)
	fmt.Println(b)
}

func copySlice(nums []int) []int {
	n := make([]int, len(nums), len(nums))
	copy(n, nums)
	return n
}
