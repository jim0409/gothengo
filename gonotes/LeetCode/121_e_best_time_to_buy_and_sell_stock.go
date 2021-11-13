package main

import "fmt"

/*
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

/*
maxProfit (預計傳入的int陣列，數字應該都是大於等於1的)
1. 將第一個數設定為最小數
2. 遇到比第一個數更小的數就回傳，更新最小的數
3. 遇到比最小數更大的數，設定為拍賣價，並且更新最佳拍賣價的拍賣日
*/
func maxProfit(prices []int) int {
	maxProfit := 0

	if len(prices) < 2 {
		return maxProfit
	}
	// sellDay 表示要賣出的日子，如果都沒有就回傳0
	sellDay := 2
	buyDay := 0
	minPrice := prices[buyDay]
	maxPrice := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
			buyDay = i
		} else {
			curcurrentProfit := prices[i] - minPrice

			if curcurrentProfit > maxProfit {
				maxProfit = curcurrentProfit
				sellDay = i
				maxPrice = prices[i]
			}
		}
	}
	fmt.Println(sellDay, maxPrice)

	return maxProfit
}

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(input))

	input2 := []int{2, 4, 1}
	fmt.Println(maxProfit(input2))

	input3 := []int{2, 1, 2, 1, 0, 1, 2}
	fmt.Println(maxProfit(input3))
}

// 參考解答
/*
func maxProfit(prices []int) int {
    max := 0
    if len(prices) == 0 {
        return max
    }

    minSoFar := prices[0]
    for _, p := range prices {
        if p < minSoFar {
            minSoFar = p
        } else {
            cur := p - minSoFar

            if cur > max {
                max = cur
            }
        }
    }

    return max
}
*/
