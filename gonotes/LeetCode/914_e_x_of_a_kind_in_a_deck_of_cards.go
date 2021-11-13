package main

/*
In a deck of cards, each card has an integer written on it.
給一疊卡片，且每一張卡片都寫上一個數字
Return true if and only if you can choose X >= 2
such that it is possible to split the entire deck into 1 or more groups of cards, where:
如果此疊卡片可以切割成多組，且每一組剛好有`X`張卡片(X >= 2)，且每組內的卡片數字皆相同，則回傳true

Each group has exactly X cards.
All the cards in each group have the same integer.

Example 1:
Input: deck = [1,2,3,4,4,3,2,1]
Output: true
Explanation: Possible partition [1,1],[2,2],[3,3],[4,4].

Example 2:
Input: deck = [1,1,1,2,2,2,3,3]
Output: false´
Explanation: No possible partition.

Example 3:
Input: deck = [1]
Output: false
Explanation: No possible partition.

Example 4:
Input: deck = [1,1]
Output: true
Explanation: Possible partition [1,1].

Example 5:
Input: deck = [1,1,2,2,2,2]
Output: true
Explanation: Possible partition [1,1],[2,2],[2,2].

Constraints:
1 <= deck.length <= 10^4
0 <= deck[i] < 10^4
*/

// 將陣列轉換組map，考慮map中每個相異數的最大公因數大於2
// https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/discuss/469539/Go-GCD-with-explanation
func hasGroupsSizeX(deck []int) bool {
	cnt := make(map[int]int, 0) // cnt := map[int]int{}
	for _, v := range deck {
		cnt[v]++
	}

	res := cnt[deck[0]]
	for _, i := range cnt {
		res = gcd(res, i)
	}

	return res >= 2
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/discuss/506935/Go-16ms-solution

func hasGroupsSizeX(deck []int) bool {
    m := map[int]int{}
    for _, v := range deck {
        m[v]++
    }
    groupSize := len(deck)
    for _, v := range m {
        if v < groupSize {
            groupSize = v
        }
    }
    if groupSize <= 1 {
        return false
    }
    for _, v := range m {
        groupSize = greatestCommonDivisor(groupSize, v)
        if groupSize <= 1 {
            return false
        }
    }
    return true
}

func greatestCommonDivisor(a, b int) int {
    l, s := a, b
    if a < b {
        l, s = b, a
    }
    for l % s != 0 {
        l, s = s, l % s
    }
    return s
}
*/
