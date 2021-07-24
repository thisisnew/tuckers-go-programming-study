package main

import "fmt"

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	var b [5]int
	b = array

	var c []int
	c = array[:]

	b[0] = 1000
	c[3] = 500

	fmt.Println("array:", array)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}
