package main

import "fmt"

type Data2 struct {
	value int
	data  [200]int
}

func ChangeData2(arg *Data2) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data2 Data2

	ChangeData2(&data2)
	fmt.Printf("value = %d\n", data2.value)
	fmt.Printf("data[100] = %d\n", data2.data[100])
}
