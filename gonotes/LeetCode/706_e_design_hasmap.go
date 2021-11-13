package main

/*
Design a HashMap without using any built-in hash table libraries.

To be specific, your design should include these functions:

*put(key, value) :
Insert a (key, value) pair into the HashMap.
If the value already exists in the HashMap, update `the value`.

*get(key):
Returns `the value` to which the specified key is mapped,
or `-1` if this map contains no mapping for the key.

*remove(key) :
Remove the mapping for the value key if this map contains the mapping for the key.


Example:
MyHashMap hashMap = new MyHashMap();
hashMap.put(1, 1);
hashMap.put(2, 2);
hashMap.get(1);            // returns 1
hashMap.get(3);            // returns -1 (not found)
hashMap.put(2, 1);          // update the existing value
hashMap.get(2);            // returns 1
hashMap.remove(2);          // remove the mapping for 2
hashMap.get(2);            // returns -1 (not found)


Note:
All keys and values will be in the range of [0, 1000000].
The number of operations will be in the range of [1, 10000].
Please do not use the built-in HashMap library.
*/

type MyHashMap struct {
	capacity int
	data     []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	capacity := 1000001
	data := make([]int, capacity)
	for i := range data {
		data[i] = -1
	}
	return MyHashMap{
		capacity: capacity,
		data:     data,
	}
}

func (hm *MyHashMap) Hash(key int) int {
	return key
	// return key % hm.capacity
}

/** value will always be non-negative. */
func (hm *MyHashMap) Put(key int, value int) {
	hm.data[hm.Hash(key)] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (hm *MyHashMap) Get(key int) int {
	return hm.data[hm.Hash(key)]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (hm *MyHashMap) Remove(key int) {
	hm.data[hm.Hash(key)] = -1
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/design-hashmap/discuss/289041/Go-100msor7.3mb-binary-tree-solution


type MyHashMap struct {
	Key, Value int
	Layer uint
	L, R *MyHashMap
}


// Initialize your data structure here.
func Constructor() MyHashMap {
    return MyHashMap{-1, -1, 0, nil, nil}
}


// value will always be non-negative.
func (this *MyHashMap) Put(key int, value int)  {
    if this.Key == -1 || this.Key == key {
		this.Key = key
		this.Value = value
	} else {
		if uint(key)&(1<<this.Layer)>0 {
			if this.L == nil {
				this.L = &MyHashMap{key, value, this.Layer+1, nil, nil}
			} else {
				this.L.Put(key, value)
			}
		} else {
			if this.R == nil {
				this.R = &MyHashMap{key, value, this.Layer+1, nil, nil}
			} else {
				this.R.Put(key, value)
			}
		}
	}
}


// Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key
func (this *MyHashMap) Get(key int) int {
    if key == this.Key {
		return this.Value
	}
	if uint(key)&(1<<this.Layer)>0 {
		if this.L != nil {
			return this.L.Get(key)
		}
	} else {
		if this.R != nil {
			return this.R.Get(key)
		}
	}
	return -1
}


// Removes the mapping of the specified value key if this map contains a mapping for the key
func (this *MyHashMap) Remove(key int)  {
    this.Put(key, -1)
}
*/

/*
solution2:
- https://leetcode.com/problems/design-hashmap/discuss/727784/Module-%2B-Array-Runtime%3A-144-ms-faster-than-54.65-of-Go-submissions

Complexity Analysis
Time: O(N / K), where n is the number of possible keys and K is total number of buckets.
Space: O(K + M), where K is the total number of buckets and M is the total number of inserted keys.



type MyHashMap struct {
    buckets [][][]int
    mod int
}


// Initialize your data structure here.
func Constructor() MyHashMap {
    var mod int = 2069

    m := MyHashMap{
        buckets: make([][][]int, mod),
        mod: mod,
    }

    return m
}


// value will always be non-negative.
func (this *MyHashMap) Put(key int, value int)  {
    hash := key % this.mod
    if len(this.buckets[hash]) == 0 {
        this.buckets[hash] = [][]int{[]int{key, value}}
    } else {
        var found bool
        for i, space := range this.buckets[hash] {
            k, _ := space[0], space[1]
            if k == key {
                this.buckets[hash][i] = []int{key, value}
                found = true
                break
            }
        }
        if !found {
            this.buckets[hash] = append(this.buckets[hash], []int{key, value})
        }
    }
}


// Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key
func (this *MyHashMap) Get(key int) int {
    hash := key % this.mod
    for _, space := range this.buckets[hash] {
        k, v := space[0], space[1]
        if k == key {
            return v
        }
    }
    return -1
}


// Removes the mapping of the specified value key if this map contains a mapping for the key
func (this *MyHashMap) Remove(key int)  {
    hash := key % this.mod
    for i, space := range this.buckets[hash] {
        k, _ := space[0], space[1]
        if k == key {
            this.buckets[hash] = append(this.buckets[hash][:i], this.buckets[hash][i+1:]...)
            break
        }
    }
}
*/
