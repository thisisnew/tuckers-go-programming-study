package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf(s.String())
}

func main() {
	s := &Student{Age: 15}
	PrintAge(s)
}
