package main

import (
	"math"
	"strconv"
)

/*
You are given a sorted unique integer array nums.

Return the smallest sorted list of ranges that cover all the numbers in the array exactly.
That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

Each range [a,b] in the list should be output as:

"a->b" if a != b
"a" if a == b


Example 1:
Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: The ranges are:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"


Example 2:
Input: nums = [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: The ranges are:
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"


Example 3:
Input: nums = []
Output: []


Example 4:
Input: nums = [-1]
Output: ["-1"]


Example 5:
Input: nums = [0]
Output: ["0"]


Constraints:

0 <= nums.length <= 20
-231 <= nums[i] <= 231 - 1
All the values of nums are unique.
nums is sorted in ascending order.
*/

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	nums = append(nums, math.MinInt32)
	last := nums[0]
	count := 0
	dest := []string{}
	for i := 1; i < len(nums); i++ {
		if nums[i] == last+count+1 {
			count++
			continue
		}
		if count == 0 {
			dest = append(dest, strconv.Itoa(last))
		} else {
			dest = append(dest, strconv.Itoa(last)+"->"+strconv.Itoa(last+count))
		}
		last = nums[i]
		count = 0
	}
	return dest
}

func main() {}

/*
refer:
- https://leetcode.com/problems/summary-ranges/discuss/1123010/Go%3A-Recursive

func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return nil
    }
    if len(nums) == 1 {
        return []string{strconv.Itoa(nums[0])}
    }
    r := make([]string, 0)
    for i:=0; i < len(nums); i++ {
        j := i+1
        for j < len(nums) && nums[j] - nums[j-1] == 1 {
            j++
        }
        if i == j-1 {
            r = append(r, strconv.Itoa(nums[i]))
            continue
        }
        r = append(r, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
        i = j-1
    }
    return r
}


- https://leetcode.com/problems/summary-ranges/discuss/788853/Go-golang-clean-solution

func summaryRanges(nums []int) []string {
    if len(nums) == 0 { return []string{} }
    ans, l, r := []string{}, 0, 0
    for i := 1; i < len(nums); i++ {
        if nums[i] - nums[i - 1] == 1 {
            r = i
        } else {
            if l == r {
                ans = append(ans, strconv.Itoa(nums[l]))
            } else {
                ans = append(ans, strconv.Itoa(nums[l]) + "->" + strconv.Itoa(nums[r]))
            }
            l, r  = i, i
        }
    }
    if l == r { return append(ans, strconv.Itoa(nums[l])) }
    return append(ans, strconv.Itoa(nums[l]) + "->" + strconv.Itoa(nums[r]))
}
*/
