package main

/*
Given an array consisting of n integers,
find the contiguous subarray of given length k that has the maximum average value.

And you need to output the maximum average value.

Example 1:
Input: [1,12,-5,-6,50,3], k = 4
Output: 12.75
Explanation:
Maximum average is (12-5-6+50)/4 = 51/4 = 12.75


Note:
1 <= k <= n <= 30,000.
Elements of the given array will be in the range [-10,000, 10,000].
*/

// 11111
// xx
//  xx
//   xx
//    xx

// func findMaxAverage(nums []int, k int) float64 {
// 	ansArray := make([]float64, len(nums)-(k-1))
// 	totoalLen := len(nums)
// 	windowSize := k
// 	initPoint := 0
// 	for i := 0; i < totoalLen-(windowSize-1); i++ {
// 		for initPtr := i; initPtr < windowSize ; intPtr ++{

// 		}
// 	}

// 	return float64(res / k)
// }

// func helper(nums []int) float64 {
// 	var res int
// 	for i := 0; i < len(nums); i++ {
// 		res = res + i
// 	}
// 	return float64(res / len(nums))
// }

func findMaxAverage(nums []int, k int) float64 {
	sum, max := 0, 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max = sum

	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/maximum-average-subarray-i/discuss/199133/go-version


func findMaxAverage(nums []int, k int) float64 {
    sum,max := 0,0
    for i:=0;i<k;i++{
        sum += nums[i]
    }
    max = sum

    for i:=k;i<len(nums);i++{
        sum = sum + nums[i] - nums[i-k]
        if sum > max{
            max = sum
        }
    }
    return float64(max)/float64(k)
}


solution2:
- https://leetcode.com/problems/maximum-average-subarray-i/discuss/407511/Go-golang-two-solutions

func findMaxAverage(nums []int, k int) float64 {

	max := math.MinInt32

	for i, j := 0, k; j <= len(nums); i, j = i+1, j+1 {
		tmp := nums[i:j]
		temp := 0
		for _, v := range tmp {
			temp += v
		}
		if temp > max {
			max = temp
		}
	}
    return float64(max)/float64(k)
}
*/
