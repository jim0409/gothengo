package main

/*
Given an integer n, return the number of trailing zeroes in n!.

Example 1:
Input: 3
Output: 0
Explanation: 3! = 6, no trailing zero.

Example 2:
Input: 5
Output: 1
Explanation: 5! = 120, one trailing zero.

Note: Your solution should be in logarithmic time complexity.
*/

func trailingZeroes(n int) int {
	res := 0
	for n != 0 {
		res += n / 5
		n /= 5
	}
	return res
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/factorial-trailing-zeroes/discuss/411416/Go-golang-0ms-solution

func trailingZeroes(n int) int {
    res := 0
    for n != 0 {
        res += n / 5
        n /= 5
    }
    return res
}
*/
