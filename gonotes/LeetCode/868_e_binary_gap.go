package main

/*
Given a positive integer n,
find and return the longest distance between any two adjacent 1's in the binary representation of n.
If there are no two adjacent 1's, return 0.

Two 1's are adjacent if there are only 0's separating them (possibly no 0's).
The distance between two 1's is the absolute difference between their bit positions.
For example, the two 1's in "1001" have a distance of 3.


Example 1:
Input: n = 22
Output: 2
Explanation: 22 in binary is "10110".
The first adjacent pair of 1's is "10110" with a distance of 2.
The second adjacent pair of 1's is "10110" with a distance of 1.
The answer is the largest of these two distances, which is 2.
Note that "10110" is not a valid pair since there is a 1 separating the two 1's underlined.

Example 2:
Input: n = 5
Output: 2
Explanation: 5 in binary is "101".

Example 3:
Input: n = 6
Output: 1
Explanation: 6 in binary is "110".

Example 4:
Input: n = 8
Output: 0
Explanation: 8 in binary is "1000".
There aren't any adjacent pairs of 1's in the binary representation of 8, so we return 0.

Example 5:
Input: n = 1
Output: 0


Constraints:

1 <= n <= 109
*/

/*
1.使用三個pointer
- 目前有1的位置
- 前一個有1的位置
- 目前有的位置

2. 宣告最大距離 dist

3. 迭代N，當N的餘數大於0時持續迭代
- digit 為N%2的餘數
- 當前位置++
	- 如果餘數為1
		前一個1的位置及為該位置
		目前有1的位置及為該位置
		- 如果前一個為1的位置為預設值`-1`
			dist = 0
		- 如果前一個為1的位置不為預設值`-1`
			dist = 當下為1的位置 減去 前一個為1的位置
	每次以N/2為新的N，持續迭代N

- 當N為0時，跳出迴圈。並且返回dist
*/
func binaryGap(N int) int {
	cur1 := -1  // current position of 1
	prev1 := -1 // previous position of 1
	pos := -1   // current position
	dist := 0   // the longest distance between two consecutive 1's

	for N > 0 {
		digit := N % 2
		pos++
		if digit == 1 {
			prev1 = cur1
			cur1 = pos
			if prev1 == -1 {
				dist = 0
			} else if cur1-prev1 > dist {
				dist = cur1 - prev1
			}
		}
		N /= 2
	}
	return dist
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/binary-gap/discuss/834404/Go-Solution-Bitwise-shifting-0ms

func binaryGap(N int) int {
	distance, max := 0, math.MinInt32

	for N > 0 {
		if N&1 == 0 {
			N = N >> 1
		} else {
			distance = 0
			N = N >> 1
			for N > 0 {
				distance++
				if N&1 == 1 {
					break
				}
				N = N >> 1
			}

			if distance > max {
				max = distance
			}
		}
	}

	return max
}
*/
