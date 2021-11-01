package main

/*
Given a non-empty integer array of size n,
find the minimum number of moves required to make all array elements equal,
where a move is incrementing n - 1 elements by 1.
每次的移動都會增加`被移動的數`增加1
透過不斷的移動，找出最小移動數。讓陣列內每個元素的值相等

Example:

Input:
[1,2,3]

Output:
3

Explanation:
Only three moves are needed (remember each move increments two elements):

[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
       1,2互換       1,2互換       1,3互換
*/

func findMin(nums []int) int {
	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

func minMoves(nums []int) int {

	diff := 0
	min := findMin(nums)

	for i := 0; i < len(nums); i++ {
		old := nums[i] - min
		diff = diff + old
	}

	return diff
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/minimum-moves-to-equal-array-elements/discuss/258633/100-use-go


func findMin(nums []int) int{
    min:=nums[0]

    for i:=1;i<len(nums);i++{
        if nums[i] < min {
            min = nums[i]
        }
    }
    return min
}

func minMoves(nums []int) int {

    diff := 0
    min := findMin(nums)

    for i:=0;i<len(nums);i++{
        old := nums[i]- min
        diff =diff+old
    }

    return diff
}
*/

/*
solution2:
- https://leetcode.com/problems/minimum-moves-to-equal-array-elements/discuss/461014/go-clean-code-32ms-beats-100


func minMoves(nums []int) int {
	sum, min := 0, math.MaxInt32
	for _, num := range nums {
		sum, min = sum+num, int(math.Min(float64(min), float64(num)))
	}
	return sum - min*len(nums)
}
*/
