package main

import "fmt"

func main() {
	a, b := "second", "first"
	fmt.Println(a, b)
	b, a = a, b
	fmt.Println(a, b)
}
