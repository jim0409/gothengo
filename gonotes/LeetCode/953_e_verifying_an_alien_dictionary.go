package main

/*
In an alien language,
surprisingly they also use english lowercase letters, but possibly in a different order.
The order of the alphabet is some permutation of lowercase letters.

Given a sequence of words written in the alien language, and the order of the alphabet,
return true if and only if the given words are sorted lexicographicaly in this alien language.


Example 1:
Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
Output: true
Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.

Example 2:
Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
Output: false
Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.

Example 3:
Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
Output: false
Explanation: The first three characters "app" match, and the second string is shorter (in size.)
According to lexicographical rules "apple" > "app", because 'l' > '∅',
where '∅' is defined as the blank character which is less than any other character (More info).


Constraints:

1 <= words.length <= 100
1 <= words[i].length <= 20
order.length == 26
All characters in words[i] and order are English lowercase letters.
*/

// https://leetcode.com/problems/verifying-an-alien-dictionary/discuss/442016/Clean-Go-solution
// 宣告一個長度26，且初始值為 0 的一個整數陣列
var pos = make([]int, 'z'-'a'+1)

func orderedWords(w1, w2 string) bool {
	for i := 0; i < len(w1) && i < len(w2); i++ {
		if w1[i] != w2[i] {
			if pos[w1[i]-'a'] > pos[w2[i]-'a'] {
				return false
			} else {
				return true
			}
		}
	}

	return len(w1) <= len(w2)
}

func isAlienSorted(words []string, order string) bool {
	for i := range order {
		pos[order[i]-'a'] = i
	}

	for i := 1; i < len(words); i++ {
		if !orderedWords(words[i-1], words[i]) {
			return false
		}
	}

	return true
}

func main() {}

/*
refer:
- https://leetcode.com/problems/verifying-an-alien-dictionary/discuss/228829/100-use-go

func isAlienSorted(words []string, order string) bool {

    var orderMp = make(map[byte] int)

    for k,v := range order {
        orderMp[byte(v)]=k
    }

    wordsLength := len(words)

    for i:=0;i<wordsLength -1;i++ {

        minWord := min(len(words[i]),len(words[i+1]))

        if len(words[i+1]) < len(words[i]) && words[i][:minWord] == words[i+1] {
            return false
        }
        for k:=0;k<minWord ;k++ {
            if orderMp[words[i][k]] == orderMp[words[i+1][k]] {
                continue
            }else if orderMp[words[i][k]] < orderMp[words[i+1][k]] {
                break
            }else {
                return false
            }
        }
    }
    return true
}

func min(x,y int) int {
    if x > y {
        return y
    }
    return x
}
*/
