/*

Value receiver :- (d Dog)

In this, Go makes a copy of the struct and gives that copy to the method.

We use it only need to read data & not to change anything.



Pointer Receiver : (d *Dog)

Here, Go passes memory address the original struct

if struct is massive, using a pointer is faster becasue Go doesn't have to spend time making copies


A good rule of thumb for beginners:

If any method of your struct needs to change data (a pointer receiver), make all the methods for that struct pointer receivers. It keeps things consistent.

*/

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
