package main

/*
Given a string, find the first non-repeating character in it and return it's index.
If it doesn't exist, return -1.

Examples1:
s = "leetcode"
return 0.

Examples2:
s = "loveleetcode",
return 2.

Note: You may assume the string contain only lowercase letters.
*/

// wrong....
// func firstUniqChar(s string) int {
// 	if len(s) == 0 {
// 		return -1
// 	}
// 	if len(s) <= 2 {
// 		if len(s) == 2 && s[0] == s[1] {
// 			return -1
// 		}
// 		return 0
// 	}
// 	m := map[rune]int{}
// 	for i, j := range s {
// 		m[j]++
// 		if m[j] > 1 {
// 			return i - 2
// 		}
// 	}
// 	return -1
// }

func firstUniqChar(s string) int {
	// 如果長度為 0 不用算既可回傳不具備重複的字元
	if len(s) == 0 {
		return -1
	}

	// 定義一個長度為 26 的整數陣列
	arr := make([]int, 26)

	// 將 s 的值依次放入 arr 中
	for _, v := range s {
		arr[byte(v)-'a']++
	}

	// 放完後，在逐項檢查，看s內的字元是否具有符合只出現一次的條件
	// 用map的方法，因為為無序的所以runtime會不穩定，但是相對的宣告 i:=0，雖然會比較穩定。但是所花費的時間是固定的~
	// for i, v := range s {
	// 	if arr[byte(v)-'a'] == 1 {
	// 		return i
	// 	}
	// }
	for i := 0; i < len(s); i++ {
		if arr[byte(s[i]-'a')] == 1 {
			return i
		}
	}

	return -1
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/first-unique-character-in-a-string/discuss/337965/Go-O(n)-Hashmap-with-explanation


func firstUniqChar(s string) int {
	d := map[byte]int{}

	// Count each character.
	for i := 0; i < len(s); i++ {
		d[s[i]]++
	}

	// Find the first unique character and return.
	for i := 0; i < len(s); i++ {
		if d[s[i]] == 1 {
			return i
		}
	}

	// If there's no unique character then return -1.
	return -1
}
*/
