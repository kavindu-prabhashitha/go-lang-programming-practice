package main

import "fmt"

type PlayerThree struct {
	health int
}

func (player *PlayerThree) takeDamageFromExplotion_three(dmg int) {
	fmt.Println("Player is taking damage from explosion")
	player.health -= dmg
}

func takeDamageFromExplotion_three(player PlayerThree, dmg int) {
	fmt.Println("Player is taking damage from explosion")
	explosionDmg := 10
	player.health -= explosionDmg
	//return
}

//Access player by pointer
func pointerExampleThree() {

	player := &PlayerThree{
		health: 100,
	}

	player02 := PlayerThree{health: 30}

	//for player 01 its pointer address
	fmt.Printf("Before explotion player  01 : %+v\n", player)
	player.takeDamageFromExplotion_three(50)
	fmt.Printf("After explotion player 01: %+v\n", player)
	println("\n")
	//player 02 is pass as value
	fmt.Printf("before explotion player 02 : %+v\n", player02)
	takeDamageFromExplotion_three(player02, 5)
	fmt.Printf("After explotion player 02: %+v\n", player02)
}

//function signatures of takeDamageFromExplotion_three is different
// takeDamageFromExplotion_three(player PlayerThree, dmg int)
// (player *PlayerThree) takeDamageFromExplotion_three(dmg int)

//When to use pointers
// If we need to do state updated
