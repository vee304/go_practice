//Writing functions and returning multiple values.

package main

import "fmt"

func calculate(a int, b int) (int, int) {
	sum := a + b
	diff := a - b

	return sum, diff
}

func main() {
	c, d := calculate(10, 4)

	fmt.Println(c, d)

}
