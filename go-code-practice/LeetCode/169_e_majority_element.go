package main

import "fmt"

/*
Given an array of size n, find the majority element.
The majority element is the element that appears more than ⌊ n/2 ⌋ times.
在一個固定的陣列中，找出現次數最高的數字

You may assume that the array is non-empty and the majority element always exist in the array.
假設陣列非空

Example 1:
Input: [3,2,3]
Output: 3

Example 2:
Input: [2,2,1,1,1,2,2]
Output: 2
*/

// func majorityElement(nums []int) int {
// 	m := map[int]int{}
// 	for _, j := range nums {
// 		m[j] = m[j] + 1
// 	}

// 	fq := 0
// 	id := 0
// 	for i, j := range m {
// 		fmt.Println(i, j)
// 		if j > fq {
// 			fq = j
// 			id = i
// 		}
// 	}
// 	fmt.Println(m)
// 	return id
// }

func majorityElement(nums []int) int {
	freq := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
		if freq[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}

	return -1
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

/*
solution:
- https://leetcode.com/problems/majority-element/discuss/341745/Go-O(n)-HashTable


func majorityElement(nums []int) int {
	freq := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
		if freq[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}

	return -1
}
*/
