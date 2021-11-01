package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Reverse(sort.IntSlice(s))
	fmt.Println(s)
}

/*
	Normally, to sort an array of integers you wrap them in an IntSlice,
	which defines the methods Len, Less, and Swap.

	These methods are in turn used by sort.Sort.
	What sort.Reverse does is that

	it takes an existing type that defines Len, Less, and Swap,
	but it replaces the Less method with a new one

	that is always the inverse of the underlying Less:
*/

// type reverse struct {
// 	// This embedded Interface permits Reverse to use the methods of
// 	// another Interface implementation.
// 	Interface
// }

// // Less returns the opposite of the embedded implementation's Less method.
// func (r reverse) Less(i, j int) bool {
// 	return r.Interface.Less(j, i)
// }

// // Reverse returns the reverse order for data.
// func Reverse(data Interface) Interface {
// 	return &reverse{data}
// }
