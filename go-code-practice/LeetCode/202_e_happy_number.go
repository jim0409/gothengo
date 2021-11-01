package main

/*
Write an algorithm to determine if a number is "happy".

A happy number is a number defined by the following process:

Starting with any positive integer,
replace the number by the sum of the squares of its digits,
and repeat the process until the number equals 1
(where it will stay),

or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy numbers.

Example:

Input: 19
Output: true
Explanation:
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1
*/

func isHappy(n int) bool {
	set := make(map[int]struct{})
	return helper(set, n)
}

// solution : https://leetcode.com/problems/happy-number/discuss/536313/Go
func helper(set map[int]struct{}, n int) bool {
	if n == 1 {
		return true
	}
	if _, ok := set[n]; ok {
		return false
	} else {
		set[n] = struct{}{}
	}

	// 算出 current sum: m
	m := 0
	for {
		if n/10 == 0 {
			m += (n % 10) * (n % 10)
			break
		} else {
			m += (n % 10) * (n % 10)
			n /= 10
		}
	}
	// 同樣對 m 做迭代
	return helper(set, m)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/happy-number/discuss/122520/4ms-Go-solution


// use a map to track used sum.
// if a sum has already been used, it's endless loop, return false,
// if sum is 1, return true.
// 創建一個 sum map，如果回傳值在sum map，表示這是一個循環數。可以停止運算，並且回傳false

func isHappy(n int) bool {
    used := make(map[int]bool)
    used[n] = true
    curr := n
    for ;curr != 1; {
        newN := 0
        for ;curr > 0; curr /= 10 {
            c:=curr%10
            newN += c*c
        }
        curr = newN
        if used[curr] {
            return false
        }
        used[curr] = true
    }
    return true
}
*/
