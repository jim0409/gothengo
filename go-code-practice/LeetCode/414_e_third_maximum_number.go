package main

import (
	"fmt"
	"math"
	"sort"
)

/*
Given a non-empty array of integers, return the third maximum number in this array.
If it does not exist, return the maximum number. The time complexity must be in O(n).

Example 1:
Input: [3, 2, 1]
Output: 1
Explanation: The third maximum is 1.


Example 2:
Input: [1, 2]
Output: 2
Explanation: The third maximum does not exist, so the maximum (2) is returned instead.

Example 3:
Input: [2, 2, 3, 1]
Output: 1
Explanation: Note that the third maximum here means the third maximum distinct number.
Both numbers with value 2 are both considered as second maximum.
*/

func thirdMax(nums []int) int {
	// 宣告一個存放整數的 map
	max := make(map[int]int)
	keys := []int{}
	for i := 0; i < len(nums); i++ {
		// 依照 nums 製作一個 max 的 map
		if _, ok := max[nums[i]]; !ok {
			max[nums[i]] = nums[i]
			/*
				將 nums[i] append 到 key 這個數字陣列
				註: 因為前面有判斷該 key 是否存在於 max 這個map內
				所以，所有的key都會是唯一值
			*/
			keys = append(keys, nums[i])
		}
	}
	// 將 keys 做排序 (小->大)
	sort.Ints(keys)

	// 如果 keys 總數小於 3 個，回傳最大值
	if len(max) < 3 {
		return keys[len(keys)-1]
	}

	// 反之，回傳第三大的值
	keys = keys[len(keys)-3 : len(keys)-1]
	return keys[0]
}

func main() {
	fmt.Println(math.MinInt64)
}

/*
solution:
- https://leetcode.com/problems/third-maximum-number/discuss/483669/go-brute-force


// 4ms 93.55% 3MB 100.00%
// brute force
// O(N) O(1)
// https://en.wikipedia.org/wiki/Brute-force_search
func thirdMax(nums []int) int {
  a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
  for _, n := range nums {
    if n <= c || n == b || n == a {
      continue
	}
	// 暴力演算法
	// 讓 a, b, c 分別為三個最小的數字
	// 首先 c = n (nums[i])
	// 如果 c > b, 將 b,c 互換 ==> b > c
	// 如果 b > a, 將 a,b 互換 ==> a > b
	// 最後結果為 a > b > c
    c = n
    if c > b {
      b, c = c, b
	}
    if b > a {
      a, b = b, a
	}
  }
  if c == math.MinInt64 {
    return a
  }
  return c
}
*/
