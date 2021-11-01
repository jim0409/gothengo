package main

/*
consecutive: 連續的
Given a binary array, find the maximum number of consecutive `1`s in this array.

Example 1:
Input: [1,1,0,1,1,1]
			  ^ ^ ^
Output: 3
Explanation:
The first two digits or the last three digits are consecutive 1s.
The maximum number of consecutive 1s is 3.

Note:
The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000
*/

func findMaxConsecutiveOnes(nums []int) int {
	length := len(nums)
	if length == 0 { // 如果長度為 0 直接回傳 0 : 沒有連續的 1
		return 0
	}

	max, cur := 0, 0 // max, cur 紀錄 目前最多的連續 1 的紀錄， cur 為一個暫存的(連續1的)紀錄
	for i := 0; i < length; i++ {
		// 如果 nums[i] 為 1，cur++
		if nums[i] == 1 {
			cur++
		} else {
			// 當nums[i] 不為 1 ，開始判斷斷點後是否需要更新 max
			if cur > max {
				max = cur
			}
			// 更新完後 cur 重置為 0
			cur = 0
		}

		// 當執行完畢最後一步時，要判斷 cur 是否大於 max
		if cur > max {
			max = cur
		}
	}
	return max
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/max-consecutive-ones/discuss/133081/Go-100


func findMaxConsecutiveOnes(a []int) int {
    best, crt := 0, 0
    for _, v := range a {
        if crt+v > best {
            best = crt+v
        }
        crt = (crt+v)*v
    }
    return best
}
*/
