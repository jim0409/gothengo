package main

import "fmt"

/*
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?


Example 1:

Input: [2,2,1]
Output: 1


Example 2:

Input: [4,1,2,1,2]
Output: 4
*/

func singleNumber(nums []int) int {

	var n int
	fmt.Println("init n is ", n)

	for _, v := range nums {
		n = n ^ v
		fmt.Println(n)
	}

	return n

}

// [[refer]] Bitwise Operators:
// https://michaelchen.tech/golang-programming/operator/

func main() {
	a := []int{1, 2, 1, 2, 4}
	fmt.Println(singleNumber(a))
}

/*
solution
https://leetcode.com/problems/single-number/discuss/357375/Go-golang-clean-solutions


solution1:

func singleNumber(nums []int) int {
    tmp := make(map[int]int)
    sum1 := 0
    sum2 := 0
    for _, v := range nums {
        sum1 += v
        tmp[v] = 1
    }
    for i := range tmp {
        sum2 += i
    }
    return 2*sum2 - sum1
}



solution2:
func singleNumber(nums []int) int {

	var n int

	for _, v := range nums {
		n = n ^ v
	}

	return n

}
*/
