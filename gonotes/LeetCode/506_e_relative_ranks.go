package main

import (
	"sort"
	"strconv"
)

/*
Given scores of N athletes, find their relative ranks and the people
with the top three highest scores, who will be awarded medals:
"Gold Medal", "Silver Medal" and "Bronze Medal".

Example 1:
Input: [5, 4, 3, 2, 1]
Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]

Explanation:
The first three athletes got the top three highest scores,
so they got "Gold Medal", "Silver Medal" and "Bronze Medal".
For the left two athletes, you just need to output their relative ranks according
to their scores.

Note:
N is a positive integer and won't exceed 10,000.
All the scores of athletes are guaranteed to be unique.
*/

func findRelativeRanks(nums []int) []string {
	index := make([][2]int, len(nums))
	for i, v := range nums {
		index[i][0], index[i][1] = v, i
	}
	sort.Slice(index, func(i, j int) bool {
		return index[i][0] > index[j][0]
	})

	ret := make([]string, len(nums))
	ret[index[0][1]] = "Gold Medal"
	if len(ret) > 1 {
		ret[index[1][1]] = "Silver Medal"
	}
	if len(ret) > 2 {
		ret[index[2][1]] = "Bronze Medal"
	}
	for i := 3; i < len(index); i++ {
		ret[index[i][1]] = strconv.Itoa(i + 1)
	}
	return ret
}

func main() {
	input := []int{5, 4, 3, 2, 1}
	findRelativeRanks(input)
}

/*
solution:
- https://leetcode.com/problems/relative-ranks/discuss/468896/go-clean-code-8ms-beats-100


func findRelativeRanks(nums []int) []string {
	athletes := make([][]int, len(nums))
	for i, num := range nums {
		athletes[i] = []int{i, num}
	}
	sort.Slice(athletes, func(i, j int) bool {
		return athletes[i][1] > athletes[j][1]
	})
	ret := make([]string, len(athletes))
	for i, athlete := range athletes {
		switch i {
		case 0:
			ret[athlete[0]] = "Gold Medal"
		case 1:
			ret[athlete[0]] = "Silver Medal"
		case 2:
			ret[athlete[0]] = "Bronze Medal"
		default:
			ret[athlete[0]] = strconv.Itoa(i + 1)
		}
	}
	return ret
}

*/
