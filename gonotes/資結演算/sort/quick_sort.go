/*
// https://zh.wikipedia.org/wiki/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F
 function quicksort(q)
 {
     var list less, pivotList, greater
     if length(q) ≤ 1
         return q
     else
     {
         select a pivot value pivot from q
         for each x in q except the pivot element
         {
             if x < pivot then add x to less
             if x ≥ pivot then add x to greater
         }
         add pivot to pivotList
         return concatenate(quicksort(less), pivotList, quicksort(greater))
     }
 }
*/

package main

import (
	"fmt"
	"math/rand"
)

func qsort(a []int) []int {
	// 如果傳入的參數a長度為2則直接回傳
	if len(a) < 2 {
		return a
	}

	// 定義前後端
	left, right := 0, len(a)-1

	// 隨機抽取一個 中間的位置 pivotIndex
	pivotIndex := rand.Int() % len(a)

	// 將 pivotIndex 與最右邊的值做交換
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// 依序迭代 a 陣列 裡面的所有值，當 a[i] < a[right] 時，將該元素與 left 位置做調換，調換之後 left++
	/*
		e.g. [1,3,5,4], pivotIndex = 1 (left:0, right:3)

		--> [1, 4, 5, 3]

		`i = 0, left = 0` =>( j=1 < a[right]=3 ) => [3,4,5,1]
		`i = 1, left = 1` =>( j=4 < a[right]=1 ) => [3,4,5,1]
		`i = 2, left = 1` =>( j=5 < a[right]=1 ) => [3,4,5,1]
		`i = 3, left = 1` =>( j=1 < a[right]=1 ) => [3,4,5,1]
	*/
	for i, j := range a {
		if j < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// 將 left right 兩位置的值做互換
	a[left], a[right] = a[right], a[left] // `left=1` & `right=3` ==> [1, 4, 5, 3]

	qsort(a[:left])   //	[]int{1}
	qsort(a[left+1:]) //	[]int{4,5,3}

	return a
}

func main() {
	a := []int{1, 3, 4, 6, 9, 6, 7, 4}
	b := qsort(a)
	fmt.Println(b)
}
