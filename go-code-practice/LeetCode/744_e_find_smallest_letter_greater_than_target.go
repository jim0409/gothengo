package main

/*
Given a list of sorted characters letters containing only lowercase letters,
and given a target letter target,

find `the smallest element` in the list that is `larger than the given target`.

Letters also wrap around. For example,
if the target is target = 'z' and letters = ['a', 'b'], the answer is 'a'.


Examples:
Input:
letters = ["c", "f", "j"]
target = "a"
Output: "c"

Input:
letters = ["c", "f", "j"]
target = "c"
Output: "f" // since `c` is equal to `c`

Input:
letters = ["c", "f", "j"]
target = "d"
Output: "f"

Input:
letters = ["c", "f", "j"]
target = "g"
Output: "j"

Input:
letters = ["c", "f", "j"]
target = "j"
Output: "c"

Input:
letters = ["c", "f", "j"]
target = "k"
Output: "c"
Note:
letters has a length in range [2, 10000].
letters consists of lowercase letters, and contains at least 2 unique letters.
target is a lowercase letter.
*/

// https://leetcode.com/problems/find-smallest-letter-greater-than-target/discuss/578972/go-clean-code-0ms-beats-100
func nextGreatestLetter(letters []byte, target byte) byte {
	i, j := 0, len(letters)
	for i < j {
		m := i + (j-i)/2
		if letters[m] <= target {
			i = m + 1
		} else {
			j = m
		}
	}
	return letters[i%len(letters)]
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/find-smallest-letter-greater-than-target/discuss/342199/Go-binary-search-solution-O(log(n))

func nextGreatestLetter(letters []byte, target byte) byte {
    idx := binarySearch(letters, target+1)
    if idx >= len(letters){
        return letters[0]
    }
    return letters[idx]
}

func binarySearch(letters []byte, target byte) int{
    hi := len(letters)
    lo := 0
    for lo < hi{
        mid := (hi+lo)/2
        if letters[mid] < target{
            lo = mid+1
        }else{
            hi = mid
        }
    }
    return lo
}
*/
