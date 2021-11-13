package main

import "bytes"

/*
International Morse Code defines a standard encoding where each letter is mapped to a series of dots and dashes,
as follows: "a" maps to ".-", "b" maps to "-...", "c" maps to "-.-.", and so on.

For convenience, the full table for the 26 letters of the English alphabet is given below:

[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
Now, given a list of words, each word can be written as a concatenation of the Morse code of each letter.

For example, "cab" can be written as "-.-..--...", (which is the concatenation "-.-." + ".-" + "-...").
We'll call such a concatenation, the transformation of a word.

Return the number of different transformations among all words we have.

Example:
Input: words = ["gin", "zen", "gig", "msg"]
Output: 2
Explanation:
The transformation of each word is:
"gin" -> "--...-."
"zen" -> "--...-."
"gig" -> "--...--."
"msg" -> "--...--."

There are 2 different transformations, "--...-." and "--...--.".

Note:
The length of words will be at most 100.
Each words[i] will have length in range [1, 12].
words[i] will only consist of lowercase letters.
*/

// func uniqueMorseRepresentations(words []string) int {

// 	// 定義 code 代表 `每個字母` 對應的 `符號`
// 	var code = map[string]string{
// 		"a": ".-",
// 		"b": "-...",
// 		"c": "-.-.",
// 		"d": "-..",
// 		"e": ".",
// 		"f": "..-.",
// 		"g": "--.",
// 		"h": "....",
// 		"i": "..",
// 		"j": ".---",
// 		"k": "-.-",
// 		"l": ".-..",
// 		"m": "--",
// 		"n": "-.",
// 		"o": "---",
// 		"p": ".--.",
// 		"q": "--.-",
// 		"r": ".-.",
// 		"s": "...",
// 		"t": "-",
// 		"u": "..-",
// 		"v": "...-",
// 		"w": ".--",
// 		"x": "-..-",
// 		"y": "-.--",
// 		"z": "--..",
// 	}

// 	m := make(map[string]int)

// 	// 轉譯字串
// 	for i := range words {
// 		var encoding string
// 		for j := range words[i] {
// 			encoding += code[string(words[i][j])]
// 		}

// 		// m[encoding] + 1
// 		m[encoding] += 1

// 	}

// 	return len(m)
// }

// - https://leetcode.com/problems/unique-morse-code-words/discuss/150272/0-ms-Go-solution
func uniqueMorseRepresentations(words []string) int {
	if len(words) == 0 {
		return 0
	}
	morse := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
		"....", "..", ".---", "-.-", ".-..", "--", "-.",
		"---", ".--.", "--.-", ".-.", "...", "-", "..-",
		"...-", ".--", "-..-", "-.--", "--..",
	}
	count := make(map[string]int)
	for _, word := range words {
		var trans bytes.Buffer
		for _, char := range word {
			trans.WriteString(morse[char-'a'])
		}
		count[trans.String()]++
	}
	return len(count)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/unique-morse-code-words/discuss/221016/100-go

func uniqueMorseRepresentations(words []string) int {
    M := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    morse_mp := make(map[string] int)

    for i:=0; i < len(words); i++ {
        morse :=""
        for _,v := range words[i] {

			// 用陣列可以省去 map 找尋的時間
            morse = morse+M[v-'a']
        }
        morse_mp[morse]++
    }
    return len(morse_mp)
}
*/
