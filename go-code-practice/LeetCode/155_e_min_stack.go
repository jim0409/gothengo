package main

import "fmt"

/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
設計一個 stack(棧) 支持 push/ pop/ top 以及 O(1) 拿取最小值

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.

Example:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.

Hint: Consider each node in the stack having a minimum value.
*/

type MinStack struct {
	Arr    []int
	MinArr []int
}

func Constructor() MinStack {
	return MinStack{
		Arr:    []int{},
		MinArr: []int{},
	}
}

func (this *MinStack) Push(x int) {
	// 如果一開始 最小陣列(MinArr) 為零，或者推入的值小於 最小陣列(MinArr) 中的最小值
	if len(this.MinArr) == 0 || x <= this.GetMin() {
		// 將輸入的 x 直接 append 到 this.MinArr 中
		this.MinArr = append(this.MinArr, x)
	}
	// 將 x 直接 append 到 Arr 中
	this.Arr = append(this.Arr, x)
}

func (this *MinStack) Pop() {
	// 如果 Arr 中的 最後一個值 等於 MinArr 中的 最後一個值
	if this.Arr[len(this.Arr)-1] == this.MinArr[len(this.MinArr)-1] {
		// 將 MinArr 的最後一個值 剔除
		this.MinArr = this.MinArr[:len(this.MinArr)-1]
	}
	// 將 Arr 中的 最後一個值 剔除
	this.Arr = this.Arr[:len(this.Arr)-1]
}

// Top 回傳 Arr 中的最後一位
func (this *MinStack) Top() int {
	return this.Arr[len(this.Arr)-1]
}

// GetMin 回傳 MinArr 中的最後一位
func (this *MinStack) GetMin() int {
	return this.MinArr[len(this.MinArr)-1]
}

func main() {
	obj := Constructor()
	obj.Push(123)
	fmt.Println(obj.Arr)
	fmt.Println(obj.GetMin())
}

/*
solution:
- https://leetcode.com/problems/min-stack/discuss/408902/Go-golang-two-solutions

type MinStack struct {
    arr []int
    min []int
}

func Constructor() MinStack {
    return MinStack{
        arr:[]int{},
        min:[]int{},
    }
}

func (this *MinStack) Push(x int)  {
    if len(this.min) == 0 || x <= this.GetMin() {
        this.min = append(this.min, x)
    }
    this.arr = append(this.arr, x)
}

func (this *MinStack) Pop()  {
    if this.arr[len(this.arr)-1] == this.min[len(this.min)-1] {
        this.min = this.min[:len(this.min)-1]
    }
    this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
    return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
    return this.min[len(this.min)-1]
}



// -- extra solution
type stack struct {
    num int
    min int
}

type MinStack struct {
    nums []stack
}

func Constructor() MinStack {
    return MinStack{
        []stack{},
    }
}

func (this *MinStack) Push(x int)  {
    if len(this.nums) == 0 {
        this.nums = append(this.nums, stack{x, x})
    } else {
        lastMin := this.nums[len(this.nums)-1].min
        if x < lastMin {
            this.nums = append(this.nums, stack{x, x})
        } else {
            this.nums = append(this.nums, stack{x, lastMin})
        }
    }
}

func (this *MinStack) Pop()  {
    this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
    return this.nums[len(this.nums)-1].num
}

func (this *MinStack) GetMin() int {
    return this.nums[len(this.nums)-1].min
}


*/
