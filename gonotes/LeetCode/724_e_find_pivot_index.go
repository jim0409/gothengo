package main

/*
Given an array of integers nums, write a method that returns the "pivot" index of this array.

We define `the pivot index` as the index where
the sum of all the numbers to the left of the index is equal to
the sum of all the numbers to the right of the index.

If no such index exists, we should return -1.
If there are multiple pivot indexes, you should return the left-most pivot index.



Example 1:
Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
The sum of the numbers to the left of index 3 (nums[3] = 6)
( > 1 + 7 + 3 =11 )
is equal to the sum of numbers to the right of index 3.
Also, 3 is the first index where this occurs.
( > 5 + 6 = 11 )

Example 2:
Input: nums = [1,2,3]
Output: -1
Explanation:
There is no index that satisfies the conditions in the problem statement.


Constraints:
The length of nums will be in the range [0, 10000].
Each element nums[i] will be an integer in the range [-1000, 1000].
*/

func pivotIndex(nums []int) int {
	var left, right int
	for _, num := range nums {
		right += num
	}
	for i, num := range nums {
		right -= num
		if left == right {
			return i
		}
		left += num
	}
	return -1
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/find-pivot-index/discuss/190614/Go-simple-and-fast-solution


func Sum(nums *[]int) int {
    sum := 0
    for _, n := range *nums {
        sum += n
    }
    return sum
}

func pivotIndex(nums []int) int {
    lsum, rsum := 0, Sum(&nums)
    for i := 0; i < len(nums); i++ {
        rsum -= nums[i]
        if rsum == lsum {
            return i
        }
        lsum += nums[i]
    }
    return -1
}
*/
