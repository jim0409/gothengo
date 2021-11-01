package main

import (
	"fmt"
	"strconv"
)

/*
https://www.itread01.com/content/1546736419.html
題意是n=1時輸出字串1；

n=2時，數上次字串中的數值個數，因為上次字串有1個1，所以輸出11；

n=3時，由於上次字元是11，有2個1，所以輸出21；

n=4時，由於上次字串是21，有1個2和1個1，所以輸出1211。依次類推，寫個countAndSay(n)函式返回字串。

The count-and-say sequence is the sequence of integers
with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.

Given an integer n where 1 ≤ n ≤ 30,
generate the nth term of the count-and-say sequence.

You can do so recursively,
in other words from the previous member read off the digits,
counting the number of digits in groups of the same digit.

Note: Each term of the sequence of integers will be represented as a string.
*/

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}
	sequence := []rune{'1'}
	for i := 2; i <= n; i++ {
		nextSequence := []rune{}
		lastRune := 'M'
		runeCount := 0
		sequence = append(sequence, 'M')
		for _, j := range sequence {
			if j != lastRune {
				if runeCount > 0 {
					nextSequence = append(nextSequence, []rune(strconv.Itoa(runeCount))...)
					nextSequence = append(nextSequence, lastRune)
				}
				lastRune = j
				runeCount = 1
				sequence = nextSequence
			} else {
				runeCount++
			}
		}
	}
	return string(sequence)
}

func main() {
	input1 := 1
	fmt.Println(countAndSay(input1))

	input2 := 4
	fmt.Println(countAndSay(input2))

}

/*
solution:
https://leetcode.com/problems/count-and-say/discuss/229646/Go-with-recursion-O(n)-%2B-question-about-go

Anybody knows how to do append(result, []byte(strconv.Itoa(length))...) in a more elegant way?

func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }

    result := handle([]byte("1"), n - 1)
    return string(result)
}

func handle(s []byte, n int) []byte {
    if n == 0 {
        return s
    }

    result, index, inputLength := []byte{}, 0, len(s)

    for index < inputLength {
        currentValue := s[index]
        lastRepeatedIndex := index

        for i := index; i < inputLength; i++ {
            if currentValue != s[i] {
                break
            }

            lastRepeatedIndex = i
        }

        length := lastRepeatedIndex - index + 1

        result = append(result, []byte(strconv.Itoa(length))...)
        result = append(result, currentValue)
        index = lastRepeatedIndex + 1
    }

    return handle(result, n-1)
}



*/
