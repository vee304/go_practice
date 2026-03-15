// 3. The Digital Pet (Multiple Field Mutations)
// Focus: Modifying multiple fields in a struct at the same time using different methods.

/* Define a struct called Pet with three fields: Name (string), Hunger (int), and Happiness (int).

Write a method called Feed. It should decrease Hunger by 20 and increase Happiness by 5. (Optional: ensure Hunger doesn't drop below 0).

Write a method called Play. It should increase Happiness by 20 but also increase Hunger by 10 (playing burns calories!).

In main, create a Pet named "Gopher" with 50 Hunger and 50 Happiness. Feed it, then play with it, and print its stats at the end using fmt.Printf.*/

package main

import "fmt"

type pet struct {
	name      string
	hunger    int
	happiness int
}

func (p *pet) Feed(hun int, hapi int) {
	p.hunger = p.hunger - hun
	p.happiness = p.happiness + hapi

	fmt.Println("current value after feeding", p.hunger, p.happiness)
}

func (p *pet) Play(hun int, hapi int) {
	p.happiness = p.happiness + hapi
	p.hunger = p.hunger + hun

	fmt.Println("current value after play", p.hunger, p.happiness)
}

func main() {
	pett := pet{
		name:      "Gojer",
		hunger:    50,
		happiness: 50,
	}

	pett.Feed(20, 5)
	pett.Play(20, 10)

}
