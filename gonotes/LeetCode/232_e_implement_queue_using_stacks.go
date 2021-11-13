package main

/*
Implement the following operations of a queue using stacks.

push(x) -- Push element x to the back of queue.
pop() -- Removes the element from in front of queue.
peek() -- Get the front element.
empty() -- Return whether the queue is empty.

Example:

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);  
queue.peek();  // returns 1
queue.pop();   // returns 1
queue.empty(); // returns false

Notes:

You must use only standard operations of a stack
which means only push to top, peek/pop from top, size, and is empty operations are valid.
*/


func main(){

}

type MyQueue struct {
	enque []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{
		enque: []int{},
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
   this.enque = append(this.enque, x) 
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	popI := this.enque[0]
	this.enque = this.enque[1:]

	return popI
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.enque[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.enque)==0
}

/*
solution:
- https://leetcode.com/problems/implement-queue-using-stacks/discuss/283479/Go-solution-using-containerlist


type MyQueue struct {
	list *list.List
}

// Initialize your data structure here.
func Constructor() MyQueue {
	l := list.New()
	return MyQueue{
		list: l,
	}
}

// Push element x to the back of queue.
func (this *MyQueue) Push(x int)  {
	this.list.PushFront(x)
}

// Removes the element from in front of queue and returns that element.
func (this *MyQueue) Pop() int {
	returnVal := this.list.Remove(this.list.Back()).(int)
	return returnVal
}

// Get the front element.
func (this *MyQueue) Peek() int {
	return this.list.Back().Value.(int)
}

// Returns whether the queue is empty.
func (this *MyQueue) Empty() bool {
	return this.list.Len() == 0
}
*/


/*
solution2:
- https://leetcode.com/problems/implement-queue-using-stacks/discuss/536289/go-one-slice-0ms-100


type MyQueue struct {
	array []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.array = append(this.array, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	top := this.Peek()
	this.array = this.array[1:len(this.array)]
	return top
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.array[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.array) == 0
}
*/