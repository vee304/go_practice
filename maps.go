/*

2. The Mini Phonebook (Maps)
Focus: Initializing maps, adding key-value pairs, and retrieving data.

The Prompt:

Create a map where the key is a string (a person's name) and the value is an integer (their age).

Add three friends and their ages to the map.

Print the age of one specific friend by looking up their name in the map.

Update one friend's age to be a year older, then print the map to verify it changed.

*/

package main

import "fmt"

func main() {

	m := map[string]int{
		"Rahul":  32,
		"Meera":  34,
		"Rakul":  35,
		"Roshan": 19,
	}

	fmt.Print("my friend Rahul's age is", m["Rahul"])

}
