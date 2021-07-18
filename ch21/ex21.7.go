package main

import "fmt"

type Writer func(string)

func writeHello(writer Writer) {
	writer("Hello World")
}

func main() {
	writeHello(func(msg string) {
		fmt.Println(msg)
	})
}
