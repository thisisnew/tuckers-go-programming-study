package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{
		balance:   100,
		firstName: "Joe",
		lastName:  "Park",
	}
	mainA.withdrawPointer(30) //포인터메서드
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20) //값타입
	fmt.Println(mainA.balance)

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance)

	mainB.withdrawPointer(30) // 포인터메서드
	fmt.Println(mainB.balance)
}
