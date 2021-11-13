package main

/*
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed,
the only constraint stopping you from robbing each of them is that
adjacent houses have security system connected and it will automatically contact the police
if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house,
determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.
*/

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		// 考慮情境
		// [1, 0, 3, 1, 1, 5] -> 1 + 3 + 5，中間的兩個 '1' 都不能選才能選到 5
		// 先預走一步，dp[i-1]+nums[i]，確認dp[i]及dp[i-1]的值哪一個比較大
		// 將比較大的那個值當作 dp[i+1]
		/*
					   i = 1, dp[2]=max(dp[1], dp[0]+nums[1]) = max(0, 1+0) => dp[2] = 1
					   i = 2, dp[3]=max(dp[2], dp[1]+nums[2]) = max(1, 0+3) => dp[3] = 3
					   i = 3, dp[4]=max(dp[3], dp[2]+nums[3]) = max(3, 1+1) => dp[4] = 3
					   i = 4, dp[5]=max(dp[4], dp[3]+nums[4]) = max(3, 3+1) => dp[5] = 4
			           i = 5, dp[6]=max(dp[5], dp[4]+nums[5]) = max(4, 3+5) => dp[6] = 8

			           可以觀察到 dp[i+1] vs dp[i-1] 恰好比較 i 前後 兩個 1 的差異
		*/
		dp[i+1] = max(dp[i], dp[i-1]+nums[i])
	}

	return dp[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/house-robber/discuss/351833/Go-DP-solution


func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i+1] = max(dp[i], dp[i-1]+nums[i])
	}
	return dp[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

/*
solution2:
- https://leetcode.com/problems/house-robber/discuss/207295/5-lines-of-GO-(almost)


func rob(nums []int) int {
    var prev1, prev2 int
    for _, num := range nums {
        prev1, prev2 = max(prev2 + num, prev1), prev1
    }

    return prev1
}

func max(n1, n2 int) int {
    if n1 > n2 {
        return n1
    }
    return n2
}
*/
