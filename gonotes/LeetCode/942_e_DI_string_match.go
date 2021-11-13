package main

/*
A permutation perm of n + 1 integers of all the integers in the range [0, n]
can be represented as a string s of length n where:

s[i] == 'I' if perm[i] < perm[i + 1], and
s[i] == 'D' if perm[i] > perm[i + 1].

Given a string s, reconstruct the permutation perm and return it.
If there are multiple valid permutations perm, return any of them.


Example 1:
Input: s = "IDID"
Output: [0,4,1,3,2]

Example 2:
Input: s = "III"
Output: [0,1,2,3]

Example 3:
Input: s = "DDI"
Output: [3,2,0,1]


Constraints:
1 <= s.length <= 105
s[i] is either 'I' or 'D'.
*/

/*
	題意: 給定一個 perm 整數數列 ，且該數列內的值皆坐落在 [0, n]
	該數列代表 字串`s` (s 為由 `I` & `D` 所構成的字串)
	如果 perm[i] < perm[i+1] 那麼 s[i] == 'I'
	反之 perm[i] > perm[i+1] 那麼 s[i] == 'D'
	假設今天給定一字串`s`，請反求該`s`字串所對應的 perm 整數數列

	1. 尋遍字串 S 所有的值
	2. 假設字元 c 為 `I` 則判斷當下的值為 0, 1, ... (小至大)
	3. 反之字元 c 為 `D` 則表示當下的值為 n, n-1, ... (大至小)
	4. 將以上結果儲存於 ret 陣列
*/
func diStringMatch(S string) []int {
	// 宣告一個回傳用的 ret 陣列
	perm := make([]int, len(S)+1)

	// 紀錄 i;小到大 & j;大到小
	i := 0
	j := len(S)

	for k, c := range S {
		if c == 'I' {
			perm[k] = i
			i++
		} else {
			perm[k] = j
			j--
		}
	}

	// 最後一個值為 i，確保符合 `I` 的性質， perm[i] < perm[i+1]
	perm[len(S)] = i

	return perm
}

func main() {}

/*
refer:
- https://leetcode.com/problems/di-string-match/discuss/508387/8ms-Go-solution-with-explanation-(Not-the-given-solution)

Start with a sorted array.
Loop through the array, for each index if it should be decreasing/increasing
and it's not, then swap the elements.

Keep doing this until you can go through the array without swapping anything.

func diStringMatch(S string) []int {
    arr := make([]int, len(S) + 1)
    for i := 0; i <= len(S); i++ {
        arr[i] = i
    }
    done := false
    for !done {
        done = true
        for i := 0; i < len(S); i++ {
            if S[i] == 'D' && arr[i] < arr[i+1] {
                arr[i], arr[i+1] = arr[i+1], arr[i]
                done = false
            } else if S[i] == 'I' && arr[i] > arr[i+1] {
                arr[i], arr[i+1] = arr[i+1], arr[i]
                done = false
            }
        }
    }
    return arr
}
*/
