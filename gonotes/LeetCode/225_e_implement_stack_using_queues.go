package main

/*
Implement the following operations of a stack using queues.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
empty() -- Return whether the stack is empty.

Example:

MyStack stack = new MyStack();

stack.push(1);
stack.push(2);  
stack.top();   // returns 2
stack.pop();   // returns 2
stack.empty(); // returns false

Notes:

You must use only standard operations of a queue -- 
which means only push to back, peek/pop from front, size, and is empty operations are valid.
You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).
*/


func main(){

}


type MyStack struct {
    enque []int // enque 作為最主要的 儲存元素陣列
    deque []int // deque 只有在做pop時會作為一個緩衝來處理 enque 移除第一個元素的需求
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{[]int{}, []int{}}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.enque = append(this.enque, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    length := len(this.enque)
    for i := 0; i < length-1; i++ {
		// deque 依序 append enque 裡面的元素
        this.deque = append(this.deque, this.enque[0])
        this.enque = this.enque[1:]
    }
    topEle := this.enque[0]
    this.enque = this.deque
    this.deque = nil
    
    return topEle
}


/** Get the top element. */
func (this *MyStack) Top() int {
    topEle := this.Pop()
    this.enque = append(this.enque, topEle)
    
    return topEle
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    if len(this.enque) == 0 {
        return true
    }
    
    return false
}




/*
solution:
- https://leetcode.com/problems/implement-stack-using-queues/discuss/264638/simple-go-solution


type MyStack struct {
	head *Node
}

type Node struct {
	value int
	next *Node
}

// Initialize your data structure here.
func Constructor() MyStack {
	return MyStack{nil}
}

// Push element x onto stack.
func (this *MyStack) Push(x int)  {
	oldHead := this.head
	this.head = &Node{x, oldHead}
}

// Removes the element on top of the stack and returns that element.
func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	currentHead := this.head
	this.head = this.head.next
	return currentHead.value
}

// Get the top element.
func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.head.value
}

// Returns whether the stack is empty.
func (this *MyStack) Empty() bool {
	return this.head == nil
}
*/