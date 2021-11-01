package main

/*
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.
Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters.
Letters are case sensitive, so "a" is considered a different type of stone from "A".

Example 1:
Input: J = "aA", S = "aAAbbbb"
Output: 3

Example 2:
Input: J = "z", S = "ZZ"
Output: 0

Note:
S and J will consist of letters and have length at most 50.
The characters in J are distinct.
*/

func numJewelsInStones(J string, S string) int {
	if len(S) == 0 || len(J) == 0 {
		return 0
	}

	// 將寶石種類map做出來，後面可以比較是否存在
	jMap := make(map[rune]bool, len(J))
	for _, j := range J {
		jMap[j] = true
	}

	// result 回傳總共符合的個數
	result := 0
	for _, s := range S {
		if jMap[s] {
			result++
		}
	}
	return result
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/jewels-and-stones/discuss/505961/go-(0ms)

func numJewelsInStones(J string, S string) int {
	result := 0

	for _, s := range S {
		for _, j := range J {
			if s == j {
				result += 1
			}
		}
	}

	return result
}
*/
