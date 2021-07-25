package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	var s Student
	var p *Student

	p = &s
	s.name = "bbb"

	fmt.Println("s:", s)
	fmt.Println("p:", p)

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
