package main

/*
You are given an array words of strings.

A move onto s consists of swapping any two even indexed characters of s, or any two odd indexed characters of s.

Two strings s and t are special-equivalent if after any number of moves onto s, s == t.

For example, s = "zzxy" and t = "xyzz" are special-equivalent because we may make the moves "zzxy" -> "xzzy" -> "xyzz" that swap s[0] and s[2], then s[1] and s[3].

Now, a group of special-equivalent strings from words is a non-empty subset of words such that:

Every pair of strings in the group are special equivalent, and;
The group is the largest size possible (ie., there isn't a string s not in the group such that s is special equivalent to every string in the group)
Return the number of groups of special-equivalent strings from words.


Example 1:
Input: words = ["abcd","cdab","cbad","xyzz","zzxy","zzyx"]
Output: 3
Explanation:
One group is ["abcd", "cdab", "cbad"], since they are all pairwise special equivalent,
and none of the other strings are all pairwise special equivalent to these.

The other two groups are ["xyzz", "zzxy"] and ["zzyx"].
> Note that in particular, "zzxy" is not special equivalent to "zzyx".

Example 2:
Input: words = ["abc","acb","bac","bca","cab","cba"]
Output: 3

Note:
1 <= words.length <= 1000
1 <= words[i].length <= 20
All words[i] have the same length.
All words[i] consist of only lowercase letters.
*/

// https://home.gamer.com.tw/artwork.php?sn=5063786
// 題意: 找尋一個字串陣列中的 特殊等價 群集數
// 特殊等價定義: 透過 `偶數次` 交換字，可以得到等同對應字的字
func numSpecialEquivGroups(A []string) int {
	// 宣告一個 map 陣列可以放置 52 個 int，此處的 int 代表的是 `a-z` 以及 `A-Z`
	// 依據 奇偶 數 將字元分開儲存
	// 奇數存放在小寫 a
	// 偶數存放在大寫 a + 26
	groupHashMap := map[[52]int]interface{}{}

	// 開始迭代字串陣列 `A`
	for _, s := range A {
		array := [52]int{}
		// 針對 `A` 裡面的每個字 `s`
		for i, c := range s {
			// 如果 i 是偶數，則將字元 array 紀錄
			if i%2 == 0 {
				array[c-'a']++
			} else {
				array[c-'a'+26]++
			}
		}
		// array 作為群族標記 groupHashMap 的一個 hash
		groupHashMap[array] = struct{}{}
	}
	return len(groupHashMap)
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/groups-of-special-equivalent-strings/discuss/581721/go-clean-code-0ms-beats-100

func numSpecialEquivGroups(A []string) int {
	set := make(map[[52]int]struct{})
	for _, word := range A {
		var data [52]int
		for i, c := range word {
			if i%2 == 0 {
				data[c-'a']++
			} else {
				data[c-'a'+26]++
			}
		}
		set[data] = struct{}{}
	}
	return len(set)
}
*/
