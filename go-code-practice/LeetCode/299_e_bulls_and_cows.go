package main

import "fmt"

/*
You are playing the following Bulls and Cows game with your friend:
You write down a number and ask your friend to guess what the number is.

Each time your friend makes a guess,
you provide a hint that indicates how many digits in said
guess match your secret number exactly in both digit and position
(called "bulls")

and how many digits match the secret number
but locate in the wrong position (called "cows").

Your friend will use successive guesses and hints
to eventually derive the secret number.


Write a function to return a hint according to the secret number
and friend's guess,
use A to indicate the bulls and B to indicate the cows.

Please note that both secret number and friend's guess may contain duplicate digits.

Example 1:

Input: secret = "1807", guess = "7810"
Output: "1A3B"
Explanation: 1 bull and 3 cows. The bull is 8, the cows are 0, 1 and 7.

Example 2:

Input: secret = "1123", guess = "0111"
Output: "1A1B"
Explanation:
The 1st 1 in friend's guess is a bull, the 2nd or 3rd 1 is a cow.

Note:
You may assume that the secret number and
your friend's guess only contain digits,
and their lengths are always equal.
*/

/*
func getHint(secret string, guess string) string {
	countA := 0
	coutnB := 0
	for i := 0; i < len(secret); i++ {
		if guess[i] == secret[i] {
			countA++
			continue
		}
		coutnB++
	}
	return fmt.Sprintf("%dA%dB", countA, coutnB)
}
*/

// runtime 0ms; Memory 2.3MB
func getHint(secret string, guess string) string {
	// 初始化，cow & bull 都是 0 次
	cow, bull := 0, 0
	// 用 count 計算 B 的情況
	count := make(map[rune]int)
	// 針對 secret 每一個 `rune` 來對 count 做存放
	for _, ch := range secret {
		count[ch]++
	}

	// 迭代 guess
	for i, ch := range guess {
		// 如果 count 中存在 ch 元素，cow++
		if count[ch] > 0 {
			cow++
			count[ch]--
		}
		// 如果 secret 的第 i 個元素剛好為 byte(ch) ; rune->byte
		// bull++ & cow-- (減去之前的 cow 計算)
		if secret[i] == byte(ch) {
			bull++
			cow--
		}
	}
	return fmt.Sprintf("%dA%dB", bull, cow)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/bulls-and-cows/discuss/340865/Go-100-runtime


// runtime 0ms; memory 2.5MB
func getHint(secret string, guess string) string {
    cowHash := map[rune]int{}
    cowCount := 0
    bullCount := 0

    secretNums := []rune(secret)
    guessNums  := []rune(guess)
    for i := 0; i < len(secret); i++ {
        if secretNums[i] != guessNums[i] {
            cowHash[secretNums[i]]++
        }
    }

    for i := 0; i < len(secret); i++ {
         if secretNums[i] == guessNums[i] {
          bullCount++
            continue
        }

        if v, ok := cowHash[guessNums[i]]; ok && v > 0 {
            cowHash[guessNums[i]]--
            cowCount++
        }
    }

    return fmt.Sprintf("%vA%vB", bullCount, cowCount)
}
*/
