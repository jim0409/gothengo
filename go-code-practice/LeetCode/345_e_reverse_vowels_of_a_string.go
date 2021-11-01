package main

/*
Write a function that takes a string as input 
and reverse only the vowels of a string.
vowels: 母音 ; a, e, i, o, u


Example 1:
Input: "hello"
Output: "holle"

Example 2:
Input: "leetcode"
Output: "leotcede"

Note:
The vowels does not include the letter "y".
*/

func reverseVowels(s string) string {
	i, j := 0, len(s)-1
	sByte := []byte(s)
	for i < j {
		if isVowel(sByte[i]) && isVowel(sByte[j]) {
			sByte[i], sByte[j] = sByte[j], sByte[i]
			i++
			j--
			continue
		}
		if !isVowel(sByte[i]) {
			i++
		}
		if !isVowel(sByte[j]) {
			j--
		}
	}
	return string(sByte)
}

func isVowel(c uint8) bool {
    if c > 'Z' {
        c = c-32
    }
	return c == 'A' || c == 'E' || c == 'I' || c == 'O' ||  c == 'U'
}


func main(){

}

/*
solution:
- https://leetcode.com/problems/reverse-vowels-of-a-string/discuss/283825/Go-Solution


func reverseVowels(s string) string {
	i, j := 0, len(s)-1
	sByte := []byte(s)
	for i < j {
		if isVowel(sByte[i]) && isVowel(sByte[j]) {
			sByte[i], sByte[j] = sByte[j], sByte[i]
			i++
			j--
			continue
		}
		if !isVowel(sByte[i]) {
			i++
		}
		if !isVowel(sByte[j]) {
			j--
		}
	}
	return string(sByte)
}

func isVowel(c uint8) bool {
    if c > 'Z' {
        c = c-32
    }
	return c == 'A' || c == 'E' || c == 'I' || c == 'O' ||  c == 'U'
}
*/