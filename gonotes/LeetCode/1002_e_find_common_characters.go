package main

/*
Given a string array words,

return an array of all characters that show up in all strings within the words

(including duplicates). You may return the answer in any order.

Example 1:
Input: words = ["bella","label","roller"]
Output: ["e","l","l"]

Example 2:
Input: words = ["cool","lock","cook"]
Output: ["c","o"]

Constraints:
1 <= words.length <= 100
1 <= words[i].length <= 100
words[i] consists of lowercase English letters.
*/

/*
wrong ...
混合再一起算例外 ..
可能，單一數組涵括多個元素，排列組合
["acabcddd","bcbdbcbd","baddbadb","cbdddcac","aacbcccd","ccccddda","cababaab","addcaccd"]

// 不能把 rune 當作 pointer，不過 map 是針對指針操作，所以也無妨 ..
// *map[rune]int invalid operation: cannot index m (variable of type *map[rune]int)
func countMap(m map[rune]int, s string) {
	for _, k := range s {
		if v, ok := m[k]; ok {
			m[k] = v + 1
		} else {
			m[k] = 1
		}
	}
}

// loop all the words and find out the len(words) or n*len(words) var
func commonChars(words []string) []string {
	l := len(words)
	// ans
	ans := []string{}
	// m 是一個 map 繼存器
	m := map[rune]int{}
	for _, w := range words {
		countMap(m, w)
	}

	for r, v := range m {
		if v >= l {
			ans = append(ans, string(r))
			if v%l == 0 {
				for i := l; i < v; i = i + l {
					ans = append(ans, string(r))
				}
			}
		}
	}

	return ans
}
*/

func commonChars(A []string) []string {
	// 26 個 英文字母
	tally := make([]int, 26)

	// 第一個字串的英文字母做 ++
	for _, ch := range A[0] {
		tally[ch-'a']++
	}

	// 從第二個英文字母開始，做多重對比
	for i := 1; i < len(A); i++ {
		// 宣告一個臨時的 t2 矩陣，計算該字的使用字母
		t2 := make([]int, 26)
		for _, ch := range A[i] {
			t2[ch-'a']++
		}

		// 比對 t2，看 t2 內是否有對應值，並且比對出現次數
		//比較次數是否小於原本統計，有的話就把 tally[i] 改為 t2
		for i, t := range tally {
			if t2[i] < t {
				tally[i] = t2[i]
			}
		}
	}

	// 比對完所有的字元後，依據 字元數目 重新組回 字元
	common := []string{}
	for ch, total := range tally {
		for total > 0 {
			total--
			common = append(common, string('a'+ch))
		}
	}
	return common
}

/*
refer:
- https://leetcode.com/problems/find-common-characters/discuss/285824/GO-O(n)-solution-beat-100-.
- https://leetcode.com/problems/find-common-characters/discuss/254619/Go-100100-Simple-Solution

func commonChars(A []string) []string {
    tmp := make([]int,26)
    for i:=range tmp{
        tmp[i]=math.MaxInt32
    }
    tmp2 := make([]int,26)
    for _,str := range A {
        for _,c := range []byte(str){
            tmp2[c-byte('a')]++
        }
        for i:=0;i<26;i++{
            tmp[i]=min(tmp[i],tmp2[i])
        }
        for i:=0;i<26;i++{
            tmp2[i]=0
        }
    }
    var rst []string
    for i,v := range tmp{
        for k:=0;k<v;k++{
            rst = append(rst,string([]byte{byte(i)+byte('a')}))
        }
    }
    return rst
}
func min(i,j int)int{
    if i<j {
        return i
    }
    return j
}
*/
