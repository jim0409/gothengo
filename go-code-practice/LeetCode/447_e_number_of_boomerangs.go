package main

/*
boomerangs: 飛旋鏢

Given n points in the plane that are all pairwise distinct, a "boomerang" is a tuple of points (i, j, k)
such that the distance between i and j equals the distance between i and k (the order of the tuple matters).
當 i,j 的距離 等同 i,k 的距離時，判斷(i,j,k)為旋鏢

Find the number of boomerangs.
You may assume that n will be at most 500 and coordinates of points are all in the range [-10000, 10000] (inclusive).

Example:

Input:
[[0,0],[1,0],[2,0]]

Output:
2

Explanation:
The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]
|1-0| = |1-2|
*/

func numberOfBoomerangs(points [][]int) int {
	var ret int
	// 將 points 中每個點拉出來看
	for _, p := range points {
		data := make(map[int]int, len(points))
		for _, q := range points {
			// 算出距離值
			d := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])

			// 如果距離已經存在 data 內，ret + 2*v
			if v, ok := data[d]; ok {
				ret += 2 * v
			}
			// 該距離值存進data後++
			data[d]++
		}
	}
	return ret
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/number-of-boomerangs/discuss/289327/Go-Solution-faster-than-100.00-less-than-100.00


func numberOfBoomerangs(points [][]int) int {
	result := 0
	for i, pointI := range points {
		hashMap := make(map[int]int, len(points))
		for j, pointJ := range points {
			if i == j {
				continue
			}
			distance := (pointJ[0]-pointI[0])*(pointJ[0]-pointI[0]) + (pointJ[1]-pointI[1])*(pointJ[1]-pointI[1])

			if n, exist := hashMap[distance]; exist {
				result += n * 2
				hashMap[distance]++
			} else {
				hashMap[distance] = 1
			}
		}
	}

	return result
}
*/
