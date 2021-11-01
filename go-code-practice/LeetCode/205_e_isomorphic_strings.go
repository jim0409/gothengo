package main

/*
Given two strings s and t, determine if they are isomorphic.
給定兩個字串，s & t，並且判斷他們是不是同構

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

Example 1:
Input: s = "egg", t = "add"
Output: true
add/ egg: 第一個字出現一次，後面兩個疊字


Example 2:
Input: s = "foo", t = "bar"
Output: false


Example 3:
Input: s = "paper", t = "title"
Output: true
paper/ title: 第一個字出現一次，後面兩個疊字

Note:
You may assume both s and t have the same length.
*/

func isIsomorphic(s string, t string) bool {
	// 如果長度不一，直接返回錯誤
	if len(s) != len(t) {
		return false
	}

	// 創作第一個 byte map 來依序儲存 s 對應位置的 byte，此處會映射到 t 對應的位置
	// 再透過 mapped 來確認該值是否存在
	mapping := map[byte]byte{}
	mapped := map[byte]interface{}{}
	for i := 0; i < len(s); i++ {
		// 從 s 開始 預設處理不存在的字串
		// 如果不存在 就對 mapping[] 賦予對應的值
		if _, exist := mapping[s[i]]; !exist {
			// 如果存在，而且 mapped[] 也存在，就表示在非結構的位置出現該值，回傳錯誤
			if _, exist := mapped[t[i]]; exist {
				return false
			}

			// 針對處理過的字串，給予 mapping[] & mapped[] 賦值
			mapping[s[i]] = t[i]
			mapped[t[i]] = nil
		} else {
			// 如果 mapping[s[i]] 出現兩次 ... 也是回傳 false
			if mapping[s[i]] != t[i] {
				return false
			}
		}
	}
	return true
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/isomorphic-strings/discuss/462980/Go-two-map


func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapping := map[byte]byte{}
	mapped := map[byte]interface{}{}
	for i := 0; i < len(s); i++ {
		if _, exist := mapping[s[i]]; !exist {
			if _, exist := mapped[t[i]]; exist {
				return false
			}
			mapping[s[i]] = t[i]
			mapped[t[i]] = nil
		} else {
			if mapping[s[i]] != t[i] {
				return false
			}
		}
	}
	return true
}
*/
