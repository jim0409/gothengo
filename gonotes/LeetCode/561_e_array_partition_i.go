package main

/*
Given an array of 2n integers, 
your task is to group these integers into n pairs of integer, 
say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of 
min(ai, bi) for all i from 1 to n as large as possible.

Example 1:
Input: [1,4,3,2]
Output: 4
Explanation: 
n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).

Note:
n is a positive integer, which is in the range of [1, 10000].
All the integers in the array will be in the range of [-10000, 10000].
*/


// 題目有點問題，他其實會需要先排序。排序好後才會按照各項元素來做最小值配對加總
// https://leetcode.com/problems/array-partition-i/discuss/497977/go-within-68ms
func arrayPairSum(nums []int) int {
    sort.Ints(nums)
    
    sum:=0
    for i:=0;i<len(nums);i=i+2{
        sum += nums[i]
    }
    
    return sum
}

func main(){

}

/*
solution:
- https://leetcode.com/problems/array-partition-i/discuss/180838/Beats-100-by-using-bucket-sort-in-Go-O(N)-52-ms

func arrayPairSum(nums []int) int {
    var buckets [20001]int
    for _, v := range nums {
        buckets[v+10000]++
    }
    sum := 0
    odd := true
    for i := 0; i < len(buckets); i++ {
        for buckets[i] > 0 {
            if odd {
                sum += i - 10000
            }
            odd = !odd
            buckets[i]--
        }
    }
    return sum
}
*/
