/*
The Prompt:
Write a program that defines three variables about a fictional character: their name (string), their age (int), and their height in meters (float64).

Use the := syntax for at least one variable.

Use the var name type syntax for at least one variable.

Print a sentence using those variables. Then, update the age (add 5 years to it) and print the new age.


*/

package main

import "fmt"

var name string = "Rishu"
var age int = 17
var height float32 = 5.3

func main() {
	fmt.Println("My name is", name, "and i'm", age, "with height of", height)
}
