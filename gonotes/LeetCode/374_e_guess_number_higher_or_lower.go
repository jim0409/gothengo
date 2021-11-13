package main

/*
We are playing the Guess Game. The game is as follows:
I pick a number from 1 to n. You have to guess which number I picked.
當使用者給定一個整數n，要從1~n這些數字中挑選一個數字來猜測。一直猜到命中數為止

Every time you guess wrong, I'll tell you whether the number is higher or lower.
You call a pre-defined API guess(int num) which returns 3 possible results (-1, 1, or 0):

-1 : My number is lower
 1 : My number is higher
 0 : Congrats! You got it!

Example :
Input: n = 10, pick = 6
Output: 6
*/

/*
  Forward declaration of guess API.
  @param  num   your guess
  @return 	     -1 if num is lower than the guess number
 			      1 if num is higher than the guess number
                otherwise return 0
  func guess(num int) int;
*/

const targetNumber = 10

func guess(num int) int {
	if num > targetNumber {
		return -1
	}
	if num < targetNumber {
		return 1
	}
	return 0
}

func guessNumber(n int) int {
	// left代表最小數，right代表最大數
	// 用二元搜尋法(binary search)來加速運算
	// https://emn178.pixnet.net/blog/post/88780407
	left, right := 1, n
	for left <= right {
		// 先猜中位數
		mid := (left + right) / 2
		res := guess(mid)
		// 猜對就返回，猜大就將 right設為`mid-1` 在求一次中位數，反之就將 left設為`mid+1`
		if res == 0 {
			return mid
		} else if res == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/guess-number-higher-or-lower/discuss/576978/go-clean-code-0-ms-beats-100

func guessNumber(n int) int {
	for i, j := 0, n; ; {
		m := i + (j-i)/2
		if ret := guess(m); ret == 1 {
			i = m + 1
		} else if ret == -1 {
			j = m - 1
		} else {
			return m
		}
	}
}
*/
