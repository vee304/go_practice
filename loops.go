//Write a program that counts from 1 to 20. If the number is even, print "[Number] is even". If it is odd, do nothing.

package main

import "fmt"

func main() {

	for i := 0; i < 20; i++ {

		if i/2 == 0 {
			fmt.Println("even number")
		} else {
			fmt.Println("odd number")
		}
	}
}
