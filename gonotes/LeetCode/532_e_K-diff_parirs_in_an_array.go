package main

/*
Given an array of integers and an integer k,
you need to find the number of unique k-diff pairs in the array.

Here a k-diff pair is defined as an integer pair (i, j),
where i and j are both numbers in the array and their absolute difference is k.


Example 1:
Input: [3, 1, 4, 1, 5], k = 2
Output: 2
Explanation:
There are two 2-diff pairs in the array, (1, 3) and (3, 5).
Although we have two 1s in the input, we should only return the number of unique pairs.

Example 2:
Input:[1, 2, 3, 4, 5], k = 1
Output: 4
Explanation:
There are four 1-diff pairs in the array, (1, 2), (2, 3), (3, 4) and (4, 5).

Example 3:
Input: [1, 3, 1, 5, 4], k = 0
Output: 1
Explanation:
There is one 0-diff pair in the array, (1, 1).

Note:
The pairs (i, j) and (j, i) count as the same pair.
The length of the array won't exceed 10,000.
All the integers in the given input belong to the range: [-1e7, 1e7].
*/

func findPairs(nums []int, k int) int {
	// 如果 差值 小於 0，直接回傳不存在
	if k < 0 {
		return 0
	}

	// data 轉換，將 nums 轉換，存成 map
	// {1, 3, 1, 5, 4} => data = {1:2, 3:1, 4:1, 5:1}
	data := make(map[int]int, len(nums))
	for _, num := range nums {
		data[num]++
	}

	var ret int // 回傳符合差值為 k 的數量
	// 當 k =0 時，只有 map 中大於 1 的 key 可以算
	if k == 0 {
		for _, v := range data {
			if v >= 2 {
				ret++
			}
		}
	} else {
		for key := range data {
			// 計算 key + k => 當下數字加上差值，均存在表示符合差值`k`
			if data[key+k] > 0 {
				ret++
			}
		}
	}

	return ret
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/k-diff-pairs-in-an-array/discuss/469025/go-clean-code-12ms-beats-100

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	data := make(map[int]int, len(nums))
	for _, num := range nums {
		data[num]++
	}
    var ret int
	if k == 0 {
		for _, v := range data {
			if v >= 2 {
				ret++
			}
		}
	} else {
		for key := range data {
			if data[key+k] > 0 {
				ret++
			}
		}
	}
	return ret
}
*/
