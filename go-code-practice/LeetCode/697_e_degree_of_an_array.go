package main

/*
Given a non-empty array of non-negative integers nums,
the degree of this array is defined as the maximum frequency of any one of its elements.

Your task is to find the smallest possible length of a (contiguous) subarray of nums,
that has the same degree as nums.

Example 1:
Input: [1, 2, 2, 3, 1]
Output: 2
Explanation:
The input array has a degree of 2 because both elements 1 and 2 appear twice.
Of the subarrays that have the same degree:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
The shortest length is 2. So return 2.

Example 2:
Input: [1,2,2,3,1,4,2]
Output: 6

Note:
nums.length will be between 1 and 50,000.
nums[i] will be an integer between 0 and 49,999.
*/

func findShortestSubArray(nums []int) int {

	tmp := make(map[int][]int)
	degree := 0
	subArr := 50000

	for i, v := range nums {
		tmp[v] = append(tmp[v], i)
		if len(tmp[v]) > degree {
			degree = len(tmp[v])
		}
	}
	for _, v := range tmp {
		if len(v) == degree {
			l := v[len(v)-1] - v[0] + 1
			if l < subArr {
				subArr = l
			}
		}
	}

	return subArr

}

func main() {

}

/*
solution:
- https://leetcode.com/problems/degree-of-an-array/discuss/269183/Better-than-Leetcode-Go-solution

func findShortestSubArray(nums []int) int {
    cache := make(map[int][]int, len(nums))
    ans := int(^uint(0) >> 1)
	maxDegree := 1
	pair := make([]int, 3)
	for i, num := range nums {
		switch v, ok := cache[num]; ok {
		case true:
			v[1] = i
			v[2]++
			pair = v
		default:
			cache[num] = append(cache[num], i, -1, 1)
			pair = cache[num]
		}

        switch {
        case pair[2] > maxDegree:
            ans = pair[1] - pair[0]
            maxDegree = pair[2]
        case pair[2] == maxDegree && pair[1] >= 0 && pair[1] - pair[0] < ans:
            ans = pair[1] - pair[0]
        }
	}
    if ans == int(^uint(0) >> 1) {
        ans = 0
    }
	return ans + 1
}
*/

/*
solution:
- https://leetcode.com/problems/degree-of-an-array/discuss/197743/go-version-beat-100

func findShortestSubArray(nums []int) int {
	nmap := make(map[int]int)
	maxnum, degree, min := 0, 0, len(nums)
	max := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		nmap[nums[i]]++
	}

	// find out degree
	for _, v := range nmap {
		if v > degree {
			degree = v
		}
	}
	// deal special case
	if degree == 1 {
		return 1
	}
	// find all candidate num
	for k, v := range nmap {
		if v == degree {
			max[maxnum] = k
			maxnum++
		}
	}
	// find min from all candidate num
	for i := 0; i < maxnum; i++ {
		current := max[i]
		CurrentDegree := degree
		first, last := -1, 0
		for j := 0; j < len(nums) && CurrentDegree > 0; j++ {
			if nums[j] == current {
				if first == -1 {
					first = j
				} else {
					last = j
				}
				CurrentDegree--
			}
		}
		CurrentLength := last - first + 1
		if CurrentLength < min {
			min = CurrentLength
		}
	}
	return min
}
*/

/*
- https://leetcode.com/problems/degree-of-an-array/solution/

Intuition and Algorithm

An array that has degree d, must have some element x occur d times.
If some subarray has the same degree, then some element x (that occured d times),
still occurs d times.

The shortest such subarray would be from the first occurrence of x until the last occurrence.

For each element in the given array,
let's know left, the index of its first occurrence;
and right, the index of its last occurrence.

For example, with nums = [1,2,3,2,5] we have left[2] = 1 and right[2] = 3.

Then, for each element x that occurs the maximum number of times,
right[x] - left[x] + 1 will be our candidate answer,
and we'll take the minimum of those candidates.
*/
