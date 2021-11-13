package main

/*
You have a `RecentCounter` class which counts the number of recent requests within a certain time frame.

Implement the RecentCounter class:

- RecentCounter() Initializes the counter with zero recent requests.
- int ping(int t) Adds a new request at time t, where t represents some time in milliseconds,
and returns the number of requests that has happened in the past 3000 milliseconds (including the new request).

Specifically, return the number of requests that have happened in the inclusive range [t - 3000, t].
It is 'guaranteed' that every call to ping uses a strictly larger value of t than the previous call.

Example 1:
Input
["RecentCounter", "ping", "ping", "ping", "ping"]
[[], [1], [100], [3001], [3002]]
Output
[null, 1, 2, 3, 3]

Explanation
RecentCounter recentCounter = new RecentCounter();
recentCounter.ping(1);     // requests = [1], range is [-2999,1], return 1
recentCounter.ping(100);   // requests = [1, 100], range is [-2900,100], return 2
recentCounter.ping(3001);  // requests = [1, 100, 3001], range is [1,3001], return 3
recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002], range is [2,3002], return 3


Constraints:

1 <= t <= 10^9
Each test case will call ping with strictly increasing values of t.
At most 10^4 calls will be made to ping.
*/

/*
	Your RecentCounter object will be instantiated and called as such:
	obj := Constructor();
	param_1 := obj.Ping(t);
*/

// https://leetcode.com/problems/number-of-recent-calls/discuss/913240/Go-Queue-solutions
// 題意：透過資料設計，設計一個方法可以算出最近 Ping 的次數，輸入當下 t 獲取範圍`(t-30000, t)`內的 Ping 次數
type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.q = append(this.q, t)
	for len(this.q) > 0 && this.q[0] < t-3000 {
		this.q = this.q[1:]
	}
	return len(this.q)
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/number-of-recent-calls/discuss/726473/Go

- Slice (as queue) implementation: (Ping() is O(N))

type RecentCounter struct {
    q []int
}


func Constructor() RecentCounter {
    return RecentCounter{[]int{}}
}


func (this *RecentCounter) Ping(t int) int {
    this.q = append(this.q, t)
    for i, val := range this.q {
        if val >= t - 3000 {
            this.q = this.q[i:]
            break
        }
    }
    return len(this.q)
}



- Binary Search: (Ping() is O(logN))

type RecentCounter struct {
    a []int
}


func Constructor() RecentCounter {
    return RecentCounter{[]int{}}
}


func (this *RecentCounter) Ping(t int) int {
    this.a = append(this.a, t)

    idx := binarySearch(this.a, t - 3000)

    return len(this.a) - idx
}

func binarySearch(arr []int, target int) int {
    lo, hi := 0, len(arr)-1

    for lo <= hi {
        mid := (lo + hi) / 2
        if arr[mid] > target {
            hi = mid - 1
        } else if arr[mid] < target {
            lo = mid + 1
        } else {
            return mid
        }
    }

    return lo
}
*/
