package main

import (
	"math"
	"sort"
)

/*
Given an integer array nums and an integer k, modify the array in the following way:

choose an index i and replace nums[i] with -nums[i].
You should apply this process exactly k times.

You may choose the same index i multiple times.

Return the largest possible sum of the array after modifying it in this way.


Example 1:
Input: nums = [4,2,3], k = 1
Output: 5
Explanation: Choose index 1 and nums becomes [4,-2,3].

Example 2:
Input: nums = [3,-1,0,2], k = 3
Output: 6
Explanation: Choose indices (1, 2, 2) and nums becomes [3,1,0,2].

Example 3:
Input: nums = [2,-3,-1,5,-4], k = 2
Output: 13
Explanation: Choose indices (1, 4) and nums becomes [2,3,-1,5,4].


Constraints:
1 <= nums.length <= 104
-100 <= nums[i] <= 100
1 <= k <= 104
*/

// https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/discuss/379972/Go-golang-clean-solution
func largestSumAfterKNegations(A []int, K int) int {

	sort.Ints(A)
	sum := 0
	min := math.MaxInt32

	for i := 0; K > 0 && i < len(A) && A[i] < 0; i++ {
		A[i] = -A[i]
		K--
	}
	for _, v := range A {
		sum += v
		if v < min {
			min = v
		}
	}

	return sum - (K%2)*min*2

}

func main() {

}

/*
refer:
- https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/discuss/854867/Go-Solution-with-sort

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	var pos int
	for K > 0 {
		if A[pos] == 0 {
			break
		}

		if pos < len(A)-1 {
			if (A[pos] < 0 && A[pos+1] < 0) || A[pos]*-1 > A[pos+1] {
				A[pos] *= -1
				pos++
			} else {
				A[pos] *= -1
			}
		} else {
			A[pos] *= -1
		}

		K--
	}

	var total int
	for _, n := range A {
		total += n
	}

	return total
}
*/
