// Code is not logically correct it's just to gauge the knowledge of interface

/*
  AN interface type provide some kind of behaviour that a type should have inorder to implement that interface

Basically an interface gives a signature a type should have to implement that interface .To say that type satisfies this interface.

*/

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64 //
}

func main() {
	var a Abser

	f := Myfloat(-math.Sqrt2)

	/*  Here we don't explicitly tell if this implements interface as long as there is a type Myfloat and it as a method whose signature and reciever matches with the interface*/

	v := Vertex{3, 4}

	a = f // Here we can assign an instance of Myfloat to a variable of type Abser interface
	a = &v

	a = v

	fmt.Println()
}

type Myfloat float64 // float type

func (f Myfloat) Abs() float64 { // signature Abs()
	fmt.Println("hfhff")
	return float64(-f)

}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(23)

}
