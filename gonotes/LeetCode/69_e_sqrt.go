package main

import (
	"fmt"
)

/*
求出一個平方根的整數
1, 4, 9, 16, 25, 36, 49, 64, 81, ...
> 1, 2, 3, 4, 5, 6, 7, 8, 9, ...

Implement int sqrt(int x).

Compute and return the square root of x,

where x is guaranteed to be a non-negative integer.

Since the return type is an integer,

the decimal digits are truncated and only the integer part of the result is returned.
*/

// solution:
// - https://leetcode.com/problems/sqrtx/discuss/338085/Go-Binary-search
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	low, high := 0, x
	var mid int
	var sqr int

	for {
		// 取中位數
		mid = (low + high) / 2
		// 如果中位數恰巧等於 最小值 直接返回 -> 收斂到 low = high = mid 了 ...
		if mid == low {
			return mid
		}

		// 計算中位數的平方，是否等同於所求值 x
		sqr = mid * mid
		if sqr == x {
			return mid
		}

		// 如果 平方 大於所求值 x，將最大值改為中位數；反之將最小值改為中位數
		if sqr > x {
			high = mid
		}
		if sqr < x {
			low = mid
		}
	}
}

func main() {
	input1 := 4
	fmt.Println(mySqrt(input1))
}

/*
refer:
- https://leetcode.com/problems/sqrtx/discuss/419298/Go-golang-solutions


solution1:
func mySqrt(x int) int {
    l := 0
    r := x
    for l < r {
        m := (r-l+1)/2 + l
        if m*m > x {
            r = m - 1
        } else {
            l = m
        }
    }
    return l
}


solution2:
func mySqrt(x int) int {
    for i := 1; i <= x; i++ {
        if i*i == x { return i }
        if i*i > x { return i-1 }
    }
    return 0
}

*/
