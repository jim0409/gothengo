package main

/*
Given a string S,
return the "reversed" string where all characters that are not a letter
stay in the same place, and all letters reverse their positions.


Example 1:
Input: "ab-cd"
Output: "dc-ba"

Example 2:
Input: "a-bC-dEf-ghIj"
Output: "j-Ih-gfE-dCba"

Example 3:
Input: "Test1ng-Leet=code-Q!"
Output: "Qedo1ct-eeLg=ntse-T!"


Note:
S.length <= 100
33 <= S[i].ASCIIcode <= 122
S doesn't contain \ or "
*/

// https://leetcode.com/problems/reverse-only-letters/discuss/282684/Go-solution-(my-go-playground-explains-rune-vs-byte)
func reverseOnlyLetters(S string) string {
	chars := []rune(S) // 將`S`轉型為[]rune{}
	n := len(chars)
	left, right := 0, n-1 // 左右兩邊同時進行

	for left < right {
		for left < n && !isLetter(chars[left]) {
			left++
		}

		for right >= 0 && !isLetter(chars[right]) {
			right--
		}

		if left < right && isLetter(chars[left]) && isLetter(chars[right]) {
			chars[left], chars[right] = chars[right], chars[left] // 需要先宣告chars，不能直接對`S`操作
			left++
			right--
		}
	}

	return string(chars)
}

// 判斷是否為字元
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/reverse-only-letters/discuss/287176/Go-Solution

func reverseOnlyLetters(S string) string {
	ss := []byte(S)

	for i, j := 0, len(ss)-1; i<j; {
		if isAlphabet1(ss[i]) && isAlphabet1(ss[j]) {
			ss[i], ss[j] = ss[j], ss[i]
			i++
			j--
			continue
		}
		if !isAlphabet1(ss[i]) && !isAlphabet1(ss[j]) {
			i++
			j--
			continue
		}
		if isAlphabet1(ss[i]) {
			j--
			continue
		}
		i++
	}

	return string(ss)
}

func isAlphabet1(r byte) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
*/
