package main

/*
Given a word, you need to judge
whether the usage of capitals in it is right or not.

We define the usage of capitals in a word to be right
when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital, like "Google".

Otherwise, we define that this word doesn't use capitals
in a right way.

Example 1:
Input: "USA"
Output: True

Example 2:
Input: "FlaG"
Output: False


Note:
The input will be a non-empty word consisting of
uppercase and lowercase latin letters.
*/

// 判斷三種情況: 1.全部小寫 2.只有第一個字大寫 3.全部大寫
func detectCapitalUse(word string) bool {
	usedCapital := false
	usedNonCapital := false

	for i := 0; i < len(word); i++ {
		if rune(word[i])-'a' >= 0 {
			usedNonCapital = true
		} else {
			if i > 0 {
				usedCapital = true
			}
		}

		if usedCapital && usedNonCapital {
			return false
		}
	}

	return true
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/detect-capital/discuss/133084/Go-1-liner

func detectCapitalUse(word string) bool {
    return strings.ToLower(word) == word || strings.ToUpper(word) == word || strings.ToLower(word[1:]) == word[1:]
}
*/
