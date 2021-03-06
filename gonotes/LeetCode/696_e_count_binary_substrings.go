package main

/*
Give a string s, count the number of non-empty (contiguous) substrings
that have the same number of 0's and 1's,
and all the 0's and all the 1's in these substrings are grouped consecutively.

Substrings that occur multiple times are counted the number of times they occur.

Example 1:
Input: "00110011"
Output: 6
Explanation:
There are 6 substrings that have equal number of consecutive 1's and 0's:
"0011", "01", "1100", "10", "0011", and "01".

Notice that some of these substrings repeat and are counted the number of times they occur.
Also, "00110011" is not a valid substring
because all the 0's (and 1's) are not grouped together.

Example 2:
Input: "10101"
Output: 4
Explanation:
There are 4 substrings: "10", "01", "10", "01"
that have equal number of consecutive 1's and 0's.

Note:
s.length will be between 1 and 50,000.
s will only consist of "0" or "1" characters.
*/

func countBinarySubstrings(s string) int {
	cur := 1
	pre := 0
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			cur++
		} else {
			res += min(cur, pre)
			pre = cur
			cur = 1
		}
	}
	return res + min(cur, pre)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/count-binary-substrings/discuss/371787/Go-golang-clean-solution-with-comment


// "011000" -> 0, 11, 000 -> [1, 2, 3] --> 取組合數
// [1, 2] -> 011 -> (01)
// [2, 3] -> 11000 -> (10, 1100)
// min(1, 2) + min(2, 3) = 3
func countBinarySubstrings(s string) int {
    cur := 1
    pre := 0
    res := 0
    for i := 1; i < len(s); i++ {
        if s[i-1] == s[i] {
            cur++
        } else {
            res += min(cur, pre)
            pre = cur
            cur = 1
        }
    }
    return res + min(cur, pre)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
*/

/*
Introduction:
- https://leetcode.com/problems/count-binary-substrings/solution/

We can convert the string s into an array groups
that represents the length of same-character contiguous blocks within the string.

For example, if s = "110001111000000", then groups = [2, 3, 4, 6].

For every binary string of the form '0' * k + '1' * k or '1' * k + '0' * k,
the middle of this string must occur between two groups.

Let's try to count the number of valid binary strings between groups[i] and groups[i+1].
If we have groups[i] = 2, groups[i+1] = 3, then it represents either "00111" or "11000".
We clearly can make min(groups[i], groups[i+1]) valid binary strings within this string.

Because the binary digits to the left or right of this string must change at the boundary,
our answer can never be larger.
*/
