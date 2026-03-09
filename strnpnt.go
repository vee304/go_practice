/*

Focus: Defining custom types (structs) and modifying them using pointers (* and &).

The Prompt:

Define a struct named Player with two fields: Name (string) and Health (int).

Write a function called TakeDamage. It should accept a pointer to a Player (*Player) and an integer representing damage. The function should subtract the damage from the player's health.

In your main function, create a Player with 100 health.

Pass the memory address of that player to your TakeDamage function (using the & operator) to hit them for 20 damage. Print the player's health before and after.


*/

package main

import "fmt"

type Player struct {
	name   string
	health int
}

func TakeDamage(p *Player, Damage int) {
	p.health = p.health - Damage
}

func main() {
	plyer := Player{
		name:   "Authur",
		health: 100,
	}
	TakeDamage(&plyer, 20)

	fmt.Println(plyer.name, plyer.health)

}
