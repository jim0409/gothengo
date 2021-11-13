package main

/*
In a string s of lowercase letters, these letters form consecutive groups of the same character.

For example, a string like s = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z", and "yy".

A group is identified by an interval [start, end],
where start and end denote the start and end indices (inclusive) of the group.

In the above example, "xxxx" has the interval [3,6].

A group is considered large if it has 3 or more characters.

Return the intervals of every large group sorted in increasing order by start index.

Example 1:
Input: s = "abbxxxxzzy"
Output: [[3,6]]
Explanation: "xxxx" is the only large group with start index 3 and end index 6.

Example 2:
Input: s = "abc"
Output: []
Explanation: We have groups "a", "b", and "c", none of which are large groups.

Example 3:
Input: s = "abcdddeeeeaabbbcd"
Output: [[3,5],[6,9],[12,14]]
Explanation: The large groups are "ddd", "eeee", and "bbb".

Example 4:
Input: s = "aba"
Output: []

Constraints:
1 <= s.length <= 1000
s contains lower-case English letters only.
*/

func largeGroupPositions(S string) [][]int {
	var rs [][]int
	var i int
	// i 預設值為 0
	for i < len(S) {
		s := i                        // 緩存 i
		b := S[i]                     // `b`為`S`字串最後一的字
		for i < len(S) && S[i] == b { // 當S[i]為字元`b`的時候，對i做計算疊加
			i++
		}

		// 如果 i-s 比 3 還要大，將i這個數添加至`rs`,起始s(之前緩存的位置),結尾`i-1`
		if i-s >= 3 {
			rs = append(rs, []int{s, i - 1})
		}
	}
	return rs
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/positions-of-large-groups/discuss/556662/Go-Two-Pointer

func largeGroupPositions(S string) [][]int {
    res := make([][]int,0)
    if len(S) <3{
        return res
    }

    i:=0
    j:=0
    for ;j<len(S);j++{
        if j == len(S)-1 || S[j] != S[j+1] {
            if j-i+1 >= 3{
                res = append(res, []int{i,j})
            }
            i = j+1
        }
    }

    return res
}
*/
