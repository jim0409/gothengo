package main

import (
	"fmt"
	"runtime"
)

var arrp chan []int

func cHeapPermutation(arr []int, size int) {
	if size == 1 {
		arrp <- arr
	}
	for i := 0; i < size; i++ {
		go cHeapPermutation(arr, size-1)

		if size%2 == 1 {
			arr[0], arr[size-1] = arr[size-1], arr[0]
		} else {
			arr[i], arr[size-1] = arr[size-1], arr[i]
		}
	}
}

func factorialPermination(size int) int {
	if size == 1 {
		return 1
	}
	return factorialPermination(size-1) * size
}

func acPermutationResult(resultNum int) [][]int {
	res := [][]int{}
	for i := 0; i < resultNum; i++ {
		select {
		case p := <-arrp:
			res = append(res, p)
		}
	}
	return res
}

func main() {
	// 平行運算的時候，因為會影響到原始的arr，所以有可能會重複運算...
	runtime.GOMAXPROCS(1)

	a := []int{1, 2, 3, 4, 5, 6}
	arrp = make(chan []int)
	go cHeapPermutation(a, len(a))

	fmt.Println(acPermutationResult(factorialPermination(len(a))))
}
