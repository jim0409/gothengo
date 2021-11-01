package main

import "strings"

/*
A sentence S is given, composed of words separated by spaces.
Each word consists of lowercase and uppercase letters only.

We would like to convert the sentence to "Goat Latin" (a made-up language similar to Pig Latin.)

The rules of Goat Latin are as follows:

If a word begins with a vowel (a, e, i, o, or u), append "ma" to the end of the word.
For example, the word 'apple' becomes 'applema'.

If a word begins with a consonant (i.e. not a vowel), remove the first letter and append it to the end, then add "ma".
For example, the word "goat" becomes "oatgma".

Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1.
For example, the first word gets "a" added to the end, the second word gets "aa" added to the end and so on.
Return the final sentence representing the conversion from S to Goat Latin.


Example 1:
Input: "I speak Goat Latin"
Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"

Example 2:
Input: "The quick brown fox jumped over the lazy dog"
Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"


Notes:
S contains only uppercase, lowercase and spaces. Exactly one space between each word.
1 <= S.length <= 150.
*/

func toGoatLatin(S string) string {
	s := strings.Fields(S)
	res := []string{}
	vowel := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	for i, word := range s {
		var firstLetter byte
		if word[0] >= 'a' && word[0] <= 'z' {
			firstLetter = word[0]
		} else {
			firstLetter = word[0] + 32
		}
		if _, ok := vowel[firstLetter]; ok {
			res = append(res, word+"ma"+strings.Repeat("a", i+1))
		} else {
			res = append(res, word[1:]+string(word[0])+"ma"+strings.Repeat("a", i+1))
		}
	}
	return strings.Join(res, " ")
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/goat-latin/discuss/255816/Go-Golang-0ms-O(n)

var vowel = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

func suffix(buf *bytes.Buffer, consonant byte, j int) {
	if consonant != 0 {
		buf.WriteByte(consonant)
	}
	buf.WriteString("ma")
	for i := 0; i < j; i++ {
		buf.WriteByte('a')
	}
}

func toGoatLatin(S string) string {
	var buf bytes.Buffer
	start := true
	var consonant byte
	j := 1
	for i := 0; i < len(S); i++ {
		c := S[i]
		if c == ' ' {
			suffix(&buf, consonant, j)
			j++
			start = true
		} else if start {
			start = false
			if !vowel[c] {
				consonant = c
				continue
			}
			consonant = 0
		}
		buf.WriteByte(c)
	}

	suffix(&buf, consonant, j)

	return buf.String()
}
*/
