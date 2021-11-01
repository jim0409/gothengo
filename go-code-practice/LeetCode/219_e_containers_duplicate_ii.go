package main

/*
Given an array of integers and an integer k,

find out whether there are two distinct indices i and j in the array 

such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

Example 1:
Input: nums = [1,2,3,1], k = 3
Output: true

Example 2:
Input: nums = [1,0,1,1], k = 1
Output: true

Example 3:
Input: nums = [1,2,3,1,2,3], k = 2
Output: false
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	// n := len(nums)
	// if n == 0 {
	// 	return false
	// }

	m := make(map[int]int)
	for i, v := range nums {
		// 如果 map裡面存在該值，比較m[v]跟對應的index位置是否小於k
		if _, ok := m[v]; ok {
			// if int(math.Abs(float64(m[v]-i))) <= k {
			if int(float64(m[v]-i)) <= k {
				return true
			} else {
				// 假設距離大於 k，則重新設定m[v]的位置為i
				if i < n {
					m[v] = i
					continue
				}
			}

		// 如果 map 裡面不存在該值，則設定該值
		} else {
			m[v] = i
		}

	}
	return false
}


func main(){

}

/*
solution:
- https://leetcode.com/problems/contains-duplicate-ii/discuss/269175/Go-golang-8-lines-solution-12ms-100


func containsNearbyDuplicate(nums []int, k int) bool {
    exist := make(map[int]int, len(nums))
    for i, v := range nums{
		// 字典回傳 key 的 pos，如果 i-pos (距離) 小於 k，則回傳 true
        if pos, ok := exist[v]; ok && i - pos <= k{
            return true
		}

		// 紀錄map[值]=pos
        exist[v] = i
    }
    return false
}
*/