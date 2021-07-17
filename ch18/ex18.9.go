package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10)
	slice3 := make([]int, 10)
	slice4 := make([]int, len(slice))

	copy(slice2, slice)
	copy(slice3, slice)
	copy(slice4, slice)

	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
}
