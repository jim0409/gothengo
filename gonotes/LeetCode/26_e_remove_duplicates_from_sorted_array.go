package main

import "fmt"

/*
Given a sorted array nums, remove the duplicates in-place

such that each element appear only once and return the new length.

Do not allocate extra space for another array,

you must do this by modifying the input array in-place with O(1) extra memory.
*/

func removeDuplicates(nums []int) int {
	m := make(map[int]int)
	for _, j := range nums {
		if m[j] == 0 {
			m[j] = 1
		} else {
			m[j] = m[j] + 1
		}
	}
	return len(m)
}

func main() {
	input := []int{1, 1, 2}
	fmt.Println(removeDuplicates(input))
	input2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(input2))
}

/*
參考解答
https://leetcode.com/problems/remove-duplicates-from-sorted-array/discuss/401418/Go-with-loop-invariants

func removeDuplicates(nums []int) int {
    // [0 ... p1] contains unique array
	// [p2 ... nums.len-1] contains new elements to be considered

	// { 1, 1, 2, 2, 4 }  --> {1, 2, 2, 2, 4} -> {1, 2, 4, 2, 4}


	p1, p2 := 0, 1

    for ; p2 < len(nums); p2++ {
        if nums[p2] != nums[p2 - 1] {
            p1++
            nums[p1] = nums[p2]
        }
    }

    return p1 + 1
}

*/
