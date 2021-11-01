package main

/*

Given two arrays, write a function to compute their intersection.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]

Note:
Each element in the result should appear
as many times as it shows in both arrays.
The result can be in any order.

Follow up:

1. What if the given array is already sorted?
How would you optimize your algorithm?

2. What if nums1's size is small compared to nums2's size?
Which algorithm is better?

3. What if elements of nums2 are stored on disk, and the memory is limited
such that you cannot load all elements into the memory at once?
*/

func intersect(nums1 []int, nums2 []int) []int {
	// 先判斷 nums1 & nums2 那一邊長度比較長
	// 如果 nums1 比較長，就將 nums1 與 nums2 做互相調換
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	// 給定一個存儲 nums1 的字典，value為該字出現的次數，預設為 0
	// 註: 這邊在宣告 c 字典時，如果帶入長度，可以在初始化就做好
	// 優點: 可以同時初始化，加速程序
	// 缺點: 後續無法調整
	c := make(map[int]int, len(nums1))

	// 如果 v 出現在字典內 1 次以上就 ++ 反之，賦予 1
	for _, v := range nums1 {
		c[v] = c[v] + 1
	}

	var inter []int

	for _, num := range nums2 {
		if val, ok := c[num]; ok && val > 0 {
			inter = append(inter, num)
			c[num]--
		}

	}
	return inter
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/intersection-of-two-arrays-ii/discuss/267296/GO-Solution-4ms


func intersect(nums1 []int, nums2 []int) []int {


    if len(nums1) > len(nums2){
        nums1, nums2 = nums2, nums1
    }


    c := make(map[int]int)

    for _, v := range nums1 {
        if _, ok := c[v]; ok{
            c[v]++
        }else{
            c[v] = 1
        }
    }

    var inter []int

    for _, num := range nums2 {
        if val, ok := c[num]; ok && val > 0{
            inter = append(inter, num)
            c[num]--
        }

    }
    return inter
}
*/
