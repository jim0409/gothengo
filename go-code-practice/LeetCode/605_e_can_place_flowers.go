package main

/*
Suppose you have a long flowerbed
in which some of the plots are planted and some are not.

However, flowers cannot be planted in adjacent plots -
they would compete for water and both would die.

Given a flowerbed
(represented as an array containing 0 and 1,
where 0 means empty and 1 means not empty),

and a number n, return if n new flowers can be planted
in it without violating the no-adjacent-flowers rule.


Example 1:
Input: flowerbed = [1,0,0,0,1], n = 1
Output: True

Example 2:
Input: flowerbed = [1,0,0,0,1], n = 2
Output: False

Note:
The input array won't violate no-adjacent-flowers rule.
The input array size is in the range of [1, 20000].
n is a non-negative integer which won't exceed the input array size.
*/

func canPlaceFlowers(flowerbed []int, n int) bool {

	// 花圃的長度 l 當作 最後邊界值
	l := len(flowerbed)
	if l == 1 && flowerbed[0] == 0 { // 如果 l == 1 而且 花圃的第一個位置為 0
		flowerbed[0] = 1 // 將第一個值替換為 1 之後 n--
		n--
	}

	if flowerbed[0] == 0 && flowerbed[1] == 0 { // 如果花圃的前兩個位置為 0 ，且長度不為 1
		flowerbed[0] = 1 // 將第一個值替換為 1 之後 n--
		n--
	}

	// 剔除 1.花圃長度為`1` 以及 2.花圃的第一、二個值為0
	// 替換，當花圃的第 i、i-1 個值為 0
	for i := 1; i < l; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
			if i == l-1 || flowerbed[i+1] == 0 { // i 為最後倒數第二個值 或 i+1 也為 0 ，f[i]設為 1 後 n--
				flowerbed[i] = 1
				n--
			}
		}
	}
	return n <= 0
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/can-place-flowers/discuss/374851/Go-Simple

````
Add a 0 at the beginning and end, then
iterate through the array and check if the curIndex - 1, curIndex, and curIndex + 1 are all 0.

If so, decrement n.

If n is <= 0 return true.
Must be less than or equal to so the case where n starts as 0 can be handled.
````

func canPlaceFlowers(flowerbed []int, n int) bool {
    flowerbed = append([]int{0}, flowerbed...)
    flowerbed = append(flowerbed, 0)
    for i := 1; i < len(flowerbed)-1; i++ {
        if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
            i++  // Not required, but might as well skip the index that doesn't need to be checked.
            n--
        }
        if n <= 0 {
            return true
        }
    }
    return false
}
*/
