package main

/*
Design a HashSet without using any built-in hash table libraries.

To be specific, your design should include these functions:

* add(value):
Insert a value into the HashSet.

*contains(value) :
Return whether the value exists in the HashSet or not.

*remove(value):
Remove a value in the HashSet.

If the value does not exist in the HashSet, do nothing.

Example:
MyHashSet hashSet = new MyHashSet();
hashSet.add(1);
hashSet.add(2);
hashSet.contains(1);    // returns true
hashSet.contains(3);    // returns false (not found)
hashSet.add(2);
hashSet.contains(2);    // returns true
hashSet.remove(2);
hashSet.contains(2);    // returns false (already removed)

Note:
All values will be in the range of [0, 1000000].
The number of operations will be in the range of [1, 10000].
Please do not use the built-in HashSet library.
*/

type MyHashSet struct {
	capacity int
	data     []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	capacity := 1000001
	return MyHashSet{
		capacity: capacity,
		data:     make([]bool, capacity),
	}
}

func (hs *MyHashSet) Hash(key int) int {
	return key // 因為此題，限制key的範圍，所以也可以不考慮生成hash func

	// 對應key，產生一組 hash_key (has_key 為 key )
	// return key % hs.capacity // could be optimize via hash key key*2^n/2^n
}

func (hs *MyHashSet) Add(key int) {
	hs.data[hs.Hash(key)] = true
}

func (hs *MyHashSet) Remove(key int) {
	hs.data[hs.Hash(key)] = false
}

/** Returns true if this set contains the specified element */
func (hs *MyHashSet) Contains(key int) bool {
	return hs.data[hs.Hash(key)]
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/design-hashset/discuss/554305/Go-chaining


const (
    mod = 1024
)

type MyHashSet struct {
    set [mod]*listNode
}

type listNode struct {
    val int
    next *listNode
}

// Initialize your data structure here.
func Constructor() MyHashSet {
    arr := [mod]*listNode{}
    return MyHashSet{set: arr}
}


func (this *MyHashSet) Add(key int)  {
    if !this.Contains(key) {
        node := &listNode{val: key, next:this.set[key % mod]}
        this.set[key % mod] = node
    }
}


func (this *MyHashSet) Remove(key int)  {
    i := key % mod
    ptr := this.set[i]
    prev := &listNode{next: ptr}
    head := prev
    for ptr != nil {
        if ptr.val == key {
            prev.next = ptr.next
            break
        } else {
            prev = prev.next
            ptr = ptr.next
        }
    }
    this.set[i] = head.next
}


// Returns true if this set contains the specified element
func (this *MyHashSet) Contains(key int) bool {
    ptr := this.set[key % mod]
    for ptr != nil {
        if ptr.val == key {
            return true
        } else {
            ptr = ptr.next
        }
    }
    return false
}
*/
