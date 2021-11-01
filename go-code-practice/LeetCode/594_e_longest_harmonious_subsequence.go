package main

import "fmt"

/*
We define a harmounious array as an array
where the difference between its maximum value and its minimum value is exactly 1.

Now, given an integer array,
// 找最大可能長度，不是絕對長度...
you need to find the length of its longest harmonious subsequence among all its possible subsequences.

Example 1:
Input: [1,3,2,2,5,2,3,7]
Output: 5
Explanation:
The longest harmonious subsequence is [3,2,2,2,3].
// 題目誤導，這邊其實是指最大長度，也就是說。拆解出來的array應該是
// [1, 3, 2, 2, 5, 2, 3, 7] => [3,2,2] + [2,3] = 5

Note: The length of the input array will not exceed 20,000.
*/

// https://www.cnblogs.com/grandyang/p/6896799.html
func findLHS(nums []int) int {

	res := 0                 // 回傳最大長度
	tmp := make(map[int]int) // 臨時map，用來儲存 nums 的 key 個數

	// loop 紀錄 nums 內所有 key 的個數
	for _, v := range nums {
		tmp[v]++
	}

	// fmt.Println(tmp) // map[1:1 2:3 3:2 5:1 7:1]

	// 把之前臨時儲存的 key 讀取出來
	for i, v := range tmp {
		// tmp 不能保證連號，所以要加上存在`ok`的判斷
		if c, ok := tmp[i+1]; ok {
			if v+c > res { // 如果 key_value + next_key_value > res，置換 res
				res = v + c
			}
		}
	}
	return res
}

func main() {
	input := []int{1, 3, 2, 2, 5, 7, 2, 3, 7}
	fmt.Println(findLHS(input))
}

/*
solution:
- https://leetcode.com/problems/longest-harmonious-subsequence/discuss/395809/Go-golang-two-solutions

func findLHS(nums []int) int {

	res := 0
	tmp := make(map[int]int)
	temp := []int{}

	for _, v := range nums {
		tmp[v]++
	}
	for i := range tmp {
		temp = append(temp, i)
	}
	sort.Ints(temp)
	for i := 1; i < len(temp); i++ {
		if temp[i]-temp[i-1] == 1 {
			if tmp[temp[i]]+tmp[temp[i-1]] > res {
				res = tmp[temp[i]] + tmp[temp[i-1]]
			}
		}
	}
	return res
}
*/

/*
solution:
- https://leetcode.com/problems/longest-harmonious-subsequence/discuss/485702/go-clean-code-48ms-beats-100

func findLHS(nums []int) int {
	data := make(map[int]int, len(nums))
	for _, num := range nums {
		data[num]++
	}
	var ret int
	for k, v := range data {
		if num, ok := data[k+1]; ok {
			ret = int(math.Max(float64(ret), float64(v+num)))
		}
	}
	return ret
}
*/
