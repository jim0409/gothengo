package main

import "fmt"

/*
A binary watch has 4 LEDs on the top which represent the hours (0-11),
and the 6 LEDs on the bottom represent the minutes (0-59).

Each LED represents a zero or one, with the least significant bit on the right.

8 4 2 1
	V V

32 16 8 4 2 1
   V  V     V

For example, the above binary watch reads "3:25".

Given a non-negative integer n which represents the number of LEDs that are currently on,
return all possible times the watch could represent.

Example:

Input: n = 1
Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]

Note:
The order of output does not matter.
The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".
The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".
*/

func readBinaryWatch(num int) []string {

	res := []string{}

	// 0 ~ 12 小時
	for h := 0; h < 12; h++ {
		// 一小時 60 分鐘
		for m := 0; m < 60; m++ {
			if helper(h)+helper(m) == num {
				s := fmt.Sprintf("%d:%02d", h, m)
				res = append(res, s)
			}
		}
	}

	return res

}

func helper(n int) int {
	res := 0
	for n > 0 {
		res += n % 2
		n /= 2
	}
	return res
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/binary-watch/discuss/391471/Go-golang-0ms-solution


func readBinaryWatch(num int) []string {

	res := []string{}

	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if helper(h)+helper(m) == num {
				s := fmt.Sprintf("%d:%02d", h, m)
				res = append(res, s)
			}
		}
	}

	return res

}

func helper(n int) int {
	res := 0
    for n > 0 {
        res += n % 2
        n /= 2
    }
	return res
}
*/
