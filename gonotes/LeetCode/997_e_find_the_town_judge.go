package man

/*
In a town, there are n people labeled from 1 to n.
There is a rumor that one of these people is secretly the town judge.

If the town judge exists, then:

The town judge trusts nobody.
Everybody (except for the town judge) trusts the town judge.
There is exactly one person that satisfies properties 1 and 2.
You are given an array trust where trust[i] = [a_i, b_i] representing that the person labeled a_i trusts the person labeled b_i.

Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.



Example 1:
Input: n = 2, trust = [[1,2]]
Output: 2

Example 2:
Input: n = 3, trust = [[1,3],[2,3]]
Output: 3

Example 3:
Input: n = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1

Example 4:
Input: n = 3, trust = [[1,2],[2,3]]
Output: -1

Example 5:
Input: n = 4, trust = [[1,3],[1,4],[2,3],[2,4],[4,3]]
Output: 3


Constraints:
- 1 <= n <= 1000
- 0 <= trust.length <= 10^4
- trust[i].length == 2
- All the pairs of trust are unique.
- a_i != b_i
- 1 <= a_i, b_i <= n
*/

// https://leetcode.com/problems/find-the-town-judge/discuss/380307/Go-golang-clean-solutions
func findJudge(N int, trust [][]int) int {
	if len(trust) == 0 {
		if N == 1 {
			return 1
		}
		return -1
	}

	people := make(map[int]int)
	judge := make(map[int]int)
	for _, v := range trust {
		people[v[0]]++
		judge[v[1]]++
	}
	for i, v := range judge {
		if _, ok := people[i]; !ok && v == N-1 {
			return i
		}
	}
	return -1
}

func main() {

}

/*
refer
- https://leetcode.com/problems/find-the-town-judge/discuss/244405/Golang-solution

func findJudge(N int, trust [][]int) int {
    m := len(trust)
    trustBalance := make([]int, N+1)

    for i := 0; i < m; i++ {
        trustBalance[trust[i][0]] -= 1
        trustBalance[trust[i][1]] += 1
    }

    for i := 1; i <= N; i++ {
        if trustBalance[i] == N-1 {
            return i
        }
    }

    return -1
}
*/
