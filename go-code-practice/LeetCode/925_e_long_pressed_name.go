package main

/*
長按錯誤
Your friend is typing his name into a keyboard.
Sometimes, when typing a character c,
the key might get long pressed, and the character will be typed 1 or more times.

You examine the typed characters of the keyboard.
Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.


Example 1:
Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.

Example 2:
Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.
e應該打兩次，但是這邊少打一次。所以打錯，就不能說長按錯誤

Example 3:
Input: name = "leelee", typed = "lleeelee"
Output: true

Example 4:
Input: name = "laiden", typed = "laiden"
Output: true
Explanation: It's not necessary to long press any character.


Constraints:
1 <= name.length <= 1000
1 <= typed.length <= 1000
The characters of name and typed are lowercase letters.
*/

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		// 判斷是否為同一字
		if name[i] != typed[j] {
			return false
		}

		// 判斷`i`小於`name`的長度，同時`j`小於`typed`的長度，且`name[i]`字母等同`typed[j]` => 考慮 name 中有重複連續字
		for ; i < len(name) && j < len(typed) && name[i] == typed[j]; i, j = i+1, j+1 {
		}

		// 考慮長按錯字，要對齊下次的字
		for ; j < len(typed) && typed[j] == typed[j-1]; j++ {
		}
	}

	// i 的長度應該恰好為 len(name)，同理 j 為 len(typed)
	return i == len(name) && j == len(typed)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/long-pressed-name/discuss/581822/go-clean-code-0ms-beats-100

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		for ; i < len(name) && j < len(typed) && name[i] == typed[j]; i, j = i+1, j+1 {}
		for ; j < len(typed) && typed[j] == typed[j-1]; j++ {}
	}
	return i == len(name) && j == len(typed)
}
*/
