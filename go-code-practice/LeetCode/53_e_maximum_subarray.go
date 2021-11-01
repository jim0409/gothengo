package main

import "fmt"

/*
Given an integer array nums,
find the contiguous subarray

(containing at least one number)
which has the largest sum and return its sum.
*/

func maxSubArray(nums []int) int {
	var res int
	var curSum int

	res = nums[0]
	if res > 0 {
		curSum = res
	} else {
		curSum = 0
	}

	for i := 1; i < len(nums); i++ {
		curSum = curSum + nums[i]

		if res < curSum {
			res = curSum
		}

		if curSum < 0 {
			curSum = 0
			// 這樣下次curSum在下次迭代的時候會變成 0 + nums[i]
			// 等於 res 直接與 curSum 做比對
		}

		// fmt.Println(curSum, res, nums[i])

	}

	return res
}

func main() {
	inputArr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4} // expect get 6, since [4, -1, 2, -1] has the largest sum 6
	fmt.Println(maxSubArray(inputArr))

	inputArr2 := []int{8, -19, 5, -4, 20}
	fmt.Println(maxSubArray(inputArr2))

	inputArr3 := []int{-1, -2}
	fmt.Println(maxSubArray(inputArr3))

	inputArr4 := []int{1, -2, 0}
	fmt.Println(maxSubArray(inputArr4))

}

/*
If you have figured out the O(n) solution,

try coding another solution using the divide and conquer approach, which is more subtle.
*/

/*
solution:
https://leetcode.com/problems/maximum-subarray/discuss/447621/Go-lang

func maxSubArray(nums []int) int {
    sum := nums[0]
    if sum < 0 {
        sum = 0
    }
    max := nums[0]
    for _,v := range nums[1:] {
        sum = sum + v
        if sum > max {
            max = sum
        }
        if sum < 0 {
            sum = 0
        }
    }
    return max
}



solution2:
https://leetcode.com/problems/maximum-subarray/discuss/403118/Go-DP-with-comments

// (1) If we know contiguous subarray with largest sum ending at any index i
// We know answer

// (2) If we know contiguous subarray with largest sum ending at i - 1
// We know (1)

// DP[i] = DP[i - 1] + nums[i] if DP[i - 1] >= 0
// DP[i] = nums[i] if DP[i - 1] < 0

func maxSubArray(nums []int) int {
    ans := nums[0]
    cur := nums[0]
    for p := 1; p < len(nums); p++ {
        if cur >= 0 {
            cur += nums[p]
        } else {
            cur = nums[p]
        }

        ans = max(ans, cur)
    }
    return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

*/
