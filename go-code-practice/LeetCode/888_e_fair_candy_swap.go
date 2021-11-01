package main

/*
Alice and Bob have candy bars of different sizes:

aliceSizes[i] is the size of the i-th bar of candy that Alice has, and
bobSizes[j] is the size of the j-th bar of candy that Bob has.

Since they are friends, they would like to exchange one candy bar each so that after the exchange,
they both have the same total amount of candy.
(The total amount of candy a person has is the sum of the sizes of candy bars they have.)

Return an integer array ans where ans[0] is the size of the candy bar that Alice must exchange,
and ans[1] is the size of the candy bar that Bob must exchange.

If there are multiple answers, you may return any one of them.
It is guaranteed an answer exists.


Example 1:
Input: aliceSizes = [1,1], bobSizes = [2,2]
Output: [1,2]

Example 2:
Input: aliceSizes = [1,2], bobSizes = [2,3]
Output: [1,2]

Example 3:
Input: aliceSizes = [2], bobSizes = [1,3]
Output: [2,3]

Example 4:
Input: aliceSizes = [1,2,5], bobSizes = [2,4]
Output: [5,4]


Note:
1 <= aliceSizes.length <= 10000
1 <= bobSizes.length <= 10000
1 <= aliceSizes[i] <= 100000
1 <= bobSizes[i] <= 100000
It is guaranteed that Alice and Bob have different total amounts of candy.
It is guaranteed there exists an answer.
*/

/*
題意，Alice與Bob是好朋友。他們各自有一定數量的糖果
給定Alice的糖果為一陣列，Bob的糖果數量亦為一陣列
兩人可以透過一次交換，讓兩人各自臃有的糖果數量相同
並將其回傳為 ans[0], ans[1] 代表兩人交換的糖果數

保證確實會存在一個 ans 為其回傳結果
*/
func fairCandySwap(A []int, B []int) []int {
	// 算出兩者糖果差距 => 後面只要找到對應差距的值，即是替換值
	diff := (sum(A) - sum(B)) / 2
	// 宣告data為一個跟 A 有相關的 map 結構
	data := make(map[int]struct{}, len(A))
	for _, num := range A {
		data[num] = struct{}{}
	}

	// 判斷 B 內的值，是否有 num+diff 在 `data` 內
	// 有則回傳
	for _, num := range B {
		if _, ok := data[num+diff]; ok {
			return []int{num + diff, num}
		}
	}
	return nil
}

func sum(nums []int) int {
	var ret int
	for _, num := range nums {
		ret += num
	}
	return ret
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/fair-candy-swap/discuss/239975/100-use-go

func sumArray(A []int) int {
    sum := 0
    for _,v := range A {
        sum = sum+ v
    }
    return sum
}
func fairCandySwap(A []int, B []int) []int {

    sumA :=sumArray(A)
    sum := sumA + sumArray(B)
    var mp = make(map[int] int)
    var ans []int
    for k,v := range B {
        mp[v] = k
    }

    for _,v := range A {
        need := sum/2 - (sumA - v)

        if _,ok := mp[need];ok {
            ans= append(ans,v,need)
            break
        }
    }
    return ans
}
*/
