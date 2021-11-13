package main

/*
On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).
在一個爬梯情境下，假設有i階梯子，每階梯子都有對應的代價。

Once you pay the cost, you can either climb one or two steps.
You need to find minimum cost to reach the top of the floor,
and you can either start from the step with index 0, or the step with index 1.
每次你都可以選擇爬一階或二階，且起始點可以從位置`0`或位置`1`開始
需要找到一個最小代價方法來爬到最後一個數字。


Example 1:
Input: cost = [10, 15, 20]
Output: 15
Explanation:
Cheapest is start on cost[1], pay that cost and go to the top.
初始位置為1(代價15)，之後爬一階即可到頂


Example 2:
Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
			   0,   1, 2, 3, 4,   5, 6, 7,   8, 9
Output: 6
Explanation:
Cheapest is start on cost[0], and only step on 1s, skipping cost[3].
初始位置為0(代價為1)，
先爬兩階，到達2(代價+1)
在爬兩階，到達4(代價+1)
在爬兩階，到達6(代價+1)
在爬一階，到達7(代價+1)
最後在爬兩階，到達9(代價+1)


Note:
cost will have a length in the range [2, 1000].
Every cost[i] will be an integer in the range [0, 999].
*/

func minCostClimbingStairs(cost []int) int {
	return withCache(cost, map[int]int{})
}

func withCache(cost []int, cache map[int]int) int {
	if len(cost) <= 1 {
		return 0
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}

	key := len(cost)
	if cache[key] != 0 {
		return cache[key]
	}

	// 比較 `爬一階` 與 `爬兩階` 的cost，選擇cost最小的。然後緩存到 cache[key]內
	cache[key] = min(cost[0]+withCache(cost[1:], cache), cost[1]+withCache(cost[2:], cache))

	return cache[key]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/min-cost-climbing-stairs/discuss/475977/Go-DP-solution


func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-2], cost[i-1])
	}
	return min(cost[len(cost)-2], cost[len(cost)-1])
}
*/
