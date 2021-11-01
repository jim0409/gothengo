package main

/*
You are playing the following Nim Game with your friend:
There is a heap of stones on the table,
each time one of you take turns to remove 1 to 3 stones.

The one who removes the last stone will be the winner.
You will take the first turn to remove the stones.

Both of you are very clever and have optimal strategies for the game.
Write a function to determine whether you can win the game
given the number of stones in the heap.

Example:

Input: 4
Output: false
Explanation:

If there are 4 stones in the heap,
then you will never win the game;
No matter 1, 2, or 3 stones you remove,
the last stone will always be removed by your friend.
*/

// 數學問題，只要取到 4 的倍數，基本上會輸
// runtime 0ms; memory 1.9MB
func canWinNim(n int) bool {
	return n%4 != 0
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/nim-game/discuss/240176/100-use-go

// runtime 0ms; memory 2MB
func canWinNim(n int) bool {
    return  (n-1) % 4 ==0||(n-2) % 4 ==0||(n-3) % 4 ==0
}
*/
