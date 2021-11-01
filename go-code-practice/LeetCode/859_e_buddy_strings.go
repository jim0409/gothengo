package main

/*
Given two strings A and B of lowercase letters,
return true if you can swap two letters in A so the result is equal to B,
otherwise, return false.

Swapping letters is defined as taking two indices i and j (0-indexed)
such that i != j and swapping the characters at A[i] and A[j].

For example, swapping at indices 0 and 2 in "abcd" results in "cbad".

Example 1:
Input: A = "ab", B = "ba"
Output: true
Explanation: You can swap A[0] = 'a' and A[1] = 'b' to get "ba", which is equal to B.

Example 2:
Input: A = "ab", B = "ab"
Output: false
Explanation: The only letters you can swap are A[0] = 'a' and A[1] = 'b', which results in "ba" != B.

Example 3:
Input: A = "aa", B = "aa"
Output: true
Explanation: You can swap A[0] = 'a' and A[1] = 'a' to get "aa", which is equal to B.

Example 4:
Input: A = "aaaaaaabc", B = "aaaaaaacb"
Output: true

Example 5:
Input: A = "", B = "aa"
Output: false


Constraints:
- 0 <= A.length <= 20000
- 0 <= B.length <= 20000
- A and B consist of lowercase letters.
*/

// https://leetcode.com/problems/buddy-strings/discuss/469727/Go-0ms-solution
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	counter := [26]int{}
	diff := [][2]byte{}
	for i := range A {
		if A[i] != B[i] {
			diff = append(diff, [2]byte{A[i], B[i]})
		} else {
			counter[A[i]-'a']++
		}
		if len(diff) > 2 {
			return false
		}
	}

	if len(diff) == 0 {
		for _, v := range counter {
			if v >= 2 {
				return true
			}
		}
		return false
	}
	if len(diff) != 2 {
		return false
	}
	if diff[0][0] == diff[1][1] && diff[0][1] == diff[1][0] {
		return true
	}
	return false
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/buddy-strings/discuss/441715/Go-solution%3A-Both-runtime-and-memory-beat-100

func buddyStrings(A string, B string) bool {
  if len(A) != len(B) {
		return false
	}
	diff := [2]int{}
	count := 0
	// store duplicate character
	dup := map[byte]int{}
	for i := 0; i < len(A); i++ {
		if (A[i] != B[i]) {
			count++
			if count <= 2 {
				diff[count - 1] = i
			} else {
				return false
			}
		}
		if _, ok := dup[A[i]]; ok {
			dup[A[i]]++
		} else {
			dup[A[i]] = 1
		}
	}
	// A equals to B
	if count == 0 {
		// check whether there has a duplicate character
		for _, v := range dup {
			if v > 1 {
				return true
			}
		}
	} else if count == 2 {
		if A[diff[0]] == B[diff[1]] && A[diff[1]] == B[diff[0]] {
			return true
		}
	}
	return false
}
*/
