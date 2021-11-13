package main

import (
	"container/heap"
	"fmt"
)

/*
Design a class to find the kth largest element in a stream.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Your KthLargest class will have a constructor,
which accepts an integer k and an integer array nums,
which contains initial elements from the stream.

For each call to the method KthLargest.add,
return the element representing the kth largest element in the stream.

Example:

int k = 3;
int[] arr = [4,5,8,2];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8

Note:
You may assume that nums' length ≥ k-1 and k ≥ 1.
*/

type KthLargest struct {
	k    int
	nums []int
}

func (this *KthLargest) Len() int {
	return len(this.nums)
}

func (this *KthLargest) Less(i, j int) bool {
	return this.nums[i] < this.nums[j]
}

func (this *KthLargest) Swap(i, j int) {
	this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
}

func (this *KthLargest) Push(x interface{}) {
	this.nums = append(this.nums, x.(int))
}

func (this *KthLargest) Pop() interface{} {
	x := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return x
}

func Constructor(k int, nums []int) KthLargest {
	obj := KthLargest{
		k:    k,
		nums: make([]int, 0, k+1),
	}
	for _, num := range nums {
		obj.Add(num)
	}
	return obj
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if len(this.nums) > this.k {
		heap.Pop(this)
	}
	return this.nums[0]
}

func main() {
	k := 3
	arr := []int{4, 5, 8, 2}

	klarget := KthLargest{
		k:    k,
		nums: arr,
	}

	fmt.Println(klarget.Add(3))
	fmt.Println(klarget.Add(5))
	fmt.Println(klarget.Add(10))
	fmt.Println(klarget.Add(9))
	fmt.Println(klarget.Add(4))
}

/*
solution:
- https://leetcode.com/problems/kth-largest-element-in-a-stream/discuss/179019/Min-Heap-implementation-in-Go


// copied from golang doc
// mininum setup of integer min heap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// real thing starts here
type KthLargest struct {
	Nums *IntHeap
	K    int
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	// push all elements to the heap
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}
	// remove the smaller elements from the heap such that only the k largest elements are in the heap
	for len(*h) > k {
		heap.Pop(h)
	}
	return KthLargest{h, k}
}

func (this *KthLargest) Add(val int) int {
	if len(*this.Nums) < this.K {
		heap.Push(this.Nums, val)
	} else if val > (*this.Nums)[0] {
		heap.Push(this.Nums, val)
		heap.Pop(this.Nums)
	}
	return (*this.Nums)[0]
}


Reference of GoLang IntHeap
https://golang.org/pkg/container/heap/
*/
