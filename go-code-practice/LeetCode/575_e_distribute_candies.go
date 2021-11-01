package main

/*
Given an integer array with even length,
where different numbers in this array represent different kinds of candies.

Each number means one candy of the corresponding kind.
You need to distribute these candies equally in number to brother and sister.

Return the maximum number of kinds of candies the sister could gain.


Example 1:
Input: candies = [1,1,2,2,3,3]
Output: 3
Explanation:
There are three different kinds of candies (1, 2 and 3), and two candies for each kind.
Optimal distribution:
The sister has candies [1,2,3] and the brother has candies [1,2,3], too.
The sister has three different kinds of candies.

Example 2:
Input: candies = [1,1,2,3]
Output: 2
Explanation:
For example, the sister has candies [2,3] and the brother has candies [1,1].
The sister has two different kinds of candies, the brother has only one kind of candies.


Note:

The length of the given array is in range [2, 10,000], and will be even.
The number in given array is in range [-100,000, 100,000].
*/

// candies 為一個有序的整數陣列，且長度為2的倍數。必須要將糖果分給brother & sister(兩類)
func distributeCandies(candies []int) int {
	ret := len(candies) / 2 // 分給 brother & sister

	// 先歸納出糖果的種類，以及每個種類的個數
	candyMap := make(map[int]int)
	for _, item := range candies {
		candyMap[item] = candyMap[item] + 1
	}

	kind := len(candyMap)

	// 如果糖果的種類個數小於 ret，那麼最大的分配數即為 kind
	if kind < ret {
		ret = kind
	}

	return ret
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/distribute-candies/discuss/289607/Go-Solution-100-100


func distributeCandies(candies []int) int {
	candyMap := make(map[int]bool, len(candies))
	for _, n := range candies {
		if !candyMap[n] {
			candyMap[n] = true
		}
	}

	average := len(candies) / 2
	if len(candyMap) <= average {
		return len(candyMap)
	}

	return average
}
*/
