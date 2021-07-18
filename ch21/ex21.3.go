package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	}

	if op == "*" {
		return mul
	}

	return nil
}

func main() {
	var operator func(int, int) int
	operator = getOperator("*")

	var result = operator(3, 4)

	fmt.Println(result)
}
