// Suppose we have a permutation containing N different elements.
// Heap found a systematic method for choosing at each step a pair of elements to switch,
// in order to produce every possible permutation of these elements exactly once.
// Let us describe Heap's method in a recursive way. First we set a counter i  to 0.
// Now we perform the following steps repeatedly until i  is equal to N.
// We use the algorithm to generate the (N − 1)! permutations of the first N − 1 elements,
// adjoining the last element to each of these. This generates all of the permutations that
// end with the last element. Then if N is odd, we switch the first element and the last one,
// while if N is even we can switch the i th element and the last one (there is no difference between
// N even and odd in the first iteration). We add one to the counter i and repeat. In each iteration,
// the algorithm will produce all of the permutations that end with the element that has just been
// moved to the "last" position. The following pseudocode outputs all permutations of a data array of length N.
// Reference https://en.wikipedia.org/wiki/Heap%27s_algorithm

package main

import "fmt"

func HeapPermutation(a []int, size int) {
	if size == 1 {
		fmt.Printf("%v ", a)
	}

	for i := 0; i < size; i++ {
		HeapPermutation(a, size-1)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

func main() {
	a := []int{1, 2, 3}
	HeapPermutation(a, len(a))
}

// # refer from : https://gist.github.com/manorie/20874a3c59e9fdfb4e184cac4130944d
