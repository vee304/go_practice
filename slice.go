/* Creating slices, using the append function, and iterating over them.
Note: In Go, Arrays have a fixed size (you can't grow them), but Slices are dynamic (they can grow). You will use Slices 99% of the time. */

package main

import "fmt"

func main() {

	ss := []string{"Banana", "Apple", "Pears", "Mangoes"} // slice without any size

	// for i := 0; i < len(ss); i++ {
	// 	fmt.Println("Vaules are as follows", ss[i])
	// }

	ss = append(ss, "Aalu", "Dragon fruit", "Kiwi", "Pineapple")

	fmt.Print(ss)

}
