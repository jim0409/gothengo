package main

import "math"

/*
Given a non-negative integer c,
your task is to decide whether there're two integers a and b such that a2 + b2 = c.

Example 1:
Input: 5
Output: True
Explanation: 1 * 1 + 2 * 2 = 5


Example 2:
Input: 3
Output: False
*/

// https://leetcode.com/problems/sum-of-square-numbers/discuss/104954/GO-0ms-2-Pointers-Solution
func judgeSquareSum(c int) bool {
	lo, hi := 0, int(math.Sqrt(float64(c)))
	for t := hi*hi + lo*lo; lo <= hi && t != c; t = hi*hi + lo*lo {
		if t < c {
			lo++
		} else {
			hi--
		}
	}

	return c == hi*hi+lo*lo
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/sum-of-square-numbers/discuss/486487/go-clean-code-0ms-beats-100

func judgeSquareSum(c int) bool {
	for i, j := 0, int(math.Sqrt(float64(c))); i <= j; {
		if sum := i*i + j*j; sum < c {
			i++
		} else if sum > c {
			j--
		} else {
			return true
		}
	}
	return false
}
*/
