package main

import "fmt"

type Animal interface {
	speak() string
}

type dog struct {
}

func (d dog) speak() string {
	return ("Woof")
}

type cat struct {
}

func (c cat) speak() string {
	return ("Meoww")
}

func makesound(a Animal) {
	fmt.Println(a.speak())
}

func main() {
	mydog := dog{}
	mycat := cat{}

	makesound(mydog)
	makesound(mycat)
}
