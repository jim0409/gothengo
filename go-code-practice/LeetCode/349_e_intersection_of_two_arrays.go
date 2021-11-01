package main

/*
Given two arrays, write a function to compute their intersection.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]

Note:
Each element in the result must be unique.
The result can be in any order.
*/

func intersection(nums1 []int, nums2 []int) []int {
	// 緩存 nums1 內的所有元素，以便後續 nums2 確認元素是否存在
	var count = map[int]bool{}
	// 將符合 交集的結果 append 到這裡
	var result = []int{}

	for _, num := range nums1 {
		count[num] = true
	}

	for _, num := range nums2 {
		if count[num] == true {
			result = append(result, num)
			count[num] = false
		}
	}
	return result
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/intersection-of-two-arrays/discuss/405483/Go-Solution


func intersection(nums1 []int, nums2 []int) []int {
    var count = map[int]bool{}
    var result = []int{}
    for _,num := range nums1{
        count[num]=true
    }
    for _, num := range nums2{
        if(count[num]==true){
            result=append(result,num)
            count[num]=false
        }
    }
    return result
}
*/
