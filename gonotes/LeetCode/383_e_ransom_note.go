package main

/*
Given an arbitrary ransom note string and another string containing letters from all the magazines,
write a function that will return true if the ransom note can be constructed from the magazines;
otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.

Note:
You may assume that both strings contain only lowercase letters.

canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true
*/

// runtime: 24ms; memory 4.3MB
func canConstruct(ransomNote string, magazine string) bool {
	// 定義一個 map 來緩存 ransomNote裡面的元素
	r := make(map[rune]int)
	for _, v := range ransomNote {
		r[v] = r[v] + 1
	}
	// 定義一個 map 來緩存 magazine裡面的元素
	m := make(map[rune]int)
	for _, v := range magazine {
		m[v] = m[v] + 1
	}

	// 比較ransomNotes內部的元素，如果該元素不存在或少於magazine map裡的元素就報錯
	for _, v := range ransomNote {
		if m[v] < r[v] {
			return false
		}
	}
	return true
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/ransom-note/discuss/510133/Simple-Go-solution


// runtime: 4ms; memory: 4.3MB
func canConstruct(ransomNote string, magazine string) bool {
	// 倘若 ransomeNote 長度大於 magazine 直接報錯
    if len(ransomNote) > len(magazine) {
        return false
	}

	// 設計一個長度為 26 的 count 陣列
	count := make([]int, 26)

	// 將 magazine 裡面每個元素 減去 a 後放入 count 的指定位置內
    for i := 0; i < len(magazine); i ++{
        count[magazine[i] - 'a'] ++
	}

	// 將 magazine 內的每個元素取出後，針對該元素將 count[...] 的對應值減 1，判斷如果結果小於 0 回傳錯誤
    for i := 0; i < len(ransomNote); i ++{
        count[ransomNote[i] - 'a'] --
        if count[ransomNote[i] - 'a'] < 0 {
            return false
        }
    }
    return true
}

// *備註*: 此題一定要先算完 magazine 內的所有元素，不能同時算。否則可能發生 m: "abaa", r: "aaab"。會報錯的狀況
*/
