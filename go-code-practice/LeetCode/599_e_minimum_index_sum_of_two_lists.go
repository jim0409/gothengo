package main

import "fmt"

/*
Suppose Andy and Doris want to choose a restaurant for dinner,
and they both have a list of favorite restaurants represented by strings.

You need to help them find out their common interest with the least list index sum.
If there is a choice tie between answers, output all of them with no order requirement.
You could assume there always exists an answer.

Example 1:
Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
Output: ["Shogun"]
Explanation:
The only restaurant they both like is "Shogun".


Example 2:
Input:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["KFC", "Shogun", "Burger King"]
Output: ["Shogun"]
Explanation:
The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).
# 雖然都有 Shogun & Burger King，但是`Shogun`的下標總和比較小，所以選擇`Shogun`

Note:
The length of both lists will be in the range of [1, 1000].
The length of strings in both lists will be in the range of [1, 30].
The index is starting from 0 to the list length minus 1.
No duplicates in both lists.
*/

// 1. compare and choose the shortest one as base list
// 2. if list2 exists the same food name as list1, add to return list
// 3. As the least common as possible
// - https://www.polarxiong.com/archives/LeetCode-599-minimum-index-sum-of-two-lists.html
// - https://leetcode.com/problems/minimum-index-sum-of-two-lists/discuss/290224/Go-Solution-beats-100
func findRestaurant(list1 []string, list2 []string) []string {
	l1Map := make(map[string]int, len(list1))
	// 將 list1 裝進 `l1Map` 裡，後面用 l1Map 來做比較基準
	for i, str := range list1 {
		l1Map[str] = i
		// fmt.Println(i) // 0, 1, 2, 3 .. 用來做 index 計算
	}

	// 回傳的 result list
	result := make([]string, 0, len(l1Map))
	index := -1

	// 依序比較 list2 內的元素
	for i, k := range list2 {
		// 如果 l1Map 存在 list2 內的菜單, j 為 l1Map 中，改菜單出現的順序
		if j, exist := l1Map[k]; exist {
			// fmt.Println(exist)
			// fmt.Println(j)
			if i+j > index && index != -1 {
				continue
			}
			if i+j == index {
				result = append(result, k)
				continue
			}
			// 選擇下標總和最小的
			index = i + j
			result = []string{k}
		}
	}

	return result
}

func main() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println(findRestaurant(list1, list2))
}

/*
solution:
- https://leetcode.com/problems/minimum-index-sum-of-two-lists/discuss/485714/go-clean-code-16ms-beats-100

func findRestaurant(list1 []string, list2 []string) []string {
	data := make(map[string]int, len(list1))
	for i, str := range list1 {
		data[str] = i
	}
	var ret []string
	min := math.MaxInt32
	for i, str := range list2 {
		v, ok := data[str]
		if !ok {
			continue
		}
		if v += i; v < min {
			ret = []string{str}
			min = v
		} else if v == min {
			ret = append(ret, str)
		}
	}
	return ret
}
*/
