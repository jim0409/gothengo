package main

/*
// https://www.cnblogs.com/grandyang/p/11318240.html
Given an array A of integers, for each integer A[i]
we may choose any x with -K <= x <= K, and add x to A[i].
給定一個陣列A，判斷裡面的每個元素A[i]
可以給定一個變數`K`，以及`x`。考慮 -K <= x <= K，

After this process, we have some array B.
Return the smallest possible difference between the maximum value of B and the minimum value of B.
假設`B`為陣列`A`加上一不定數`x`後的產物。
找出在陣列`B`中的最小公差值

==> 如何讓兩者之間的差值最小? 讓最大值盡可能變小，最小值盡可能變大 ==> (max-K)-(min+K)

Example 1:
Input: A = [1], K = 0
Output: 0
Explanation: B = [1]

Example 2:
Input: A = [0,10], K = 2
Output: 6
Explanation: B = [2,8]

Example 3:
Input: A = [1,3,6], K = 3
Output: 0
Explanation: B = [3,3,3] or B = [4,4,4]


Note:
1 <= A.length <= 10000
0 <= A[i] <= 10000
0 <= K <= 10000
*/

// https://leetcode.com/problems/smallest-range-i/discuss/182186/Go-Beats-100-without-any-helper-time-O(N)-20ms
func smallestRangeI(A []int, K int) int {
	min, max := A[0], A[0]
	for _, v := range A {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	ret := (max - K) - (min + K)
	if ret < 0 {
		return 0
	}
	return ret
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/smallest-range-i/discuss/581767/go-clean-code-12ms-beats-100

func smallestRangeI(A []int, K int) int {
	max, min := math.MinInt32, math.MaxInt32
	for _, num := range A {
		max, min = int(math.Max(float64(max), float64(num))), int(math.Min(float64(min), float64(num)))
	}
	return int(math.Max(float64(0), float64(max-min-2*K)))
}
*/
