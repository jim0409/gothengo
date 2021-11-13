package main

/*
You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.
Given n, find the total number of full staircase rows that can be formed.
n is a non-negative integer and fits within the range of a 32-bit signed integer.

給n個硬幣，用這n個硬幣來排列出階梯，求最多可以排幾階(備註:階梯必須要完整)

Example 1:

n = 5
The coins can form the following rows:
*
**
**

Because the 3rd row is incomplete, we return 2.

Example 2:
n = 8
The coins can form the following rows:
*
**
***
**

Because the 4th row is incomplete, we return 3.
*/

func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}
	count := 1
	for n > 0 {
		n = n - count
		if n-count > 0 {
			count++
		}
	}
	return count
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/arranging-coins/discuss/92275/Go-Newton-Iterative-Sqrt-Method


// 0ms 100%; 2.2MB 100%
func arrangeCoins(n int) int {
	n *= 2
	x := n
	for !(x*(x+1) <= n && n < (x+1)*(x+2)) {
		x = (x + n/(x+1)) / 2
	}
	return x
}
*/

/*
solution2:
- https://leetcode.com/problems/arranging-coins/discuss/412913/Go-lang-Solution


func arrangeCoins(n int) int {
    cnt:=1
    for n>=cnt{
        n-=cnt
        cnt++
    }
    return cnt-1
}
*/
