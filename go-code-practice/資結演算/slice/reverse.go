package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	reverse(a, 0, len(a)-1)
	fmt.Println(a) //[7 6 5 4 3 2 1]
}

func reverse(seq []int, i, j int) {
	for i < j {
		seq[i], seq[j] = seq[j], seq[i]
		i++
		j--
	}
}
