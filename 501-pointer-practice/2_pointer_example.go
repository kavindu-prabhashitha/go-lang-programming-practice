package main

import "fmt"

type PlayerTwo struct {
	health int
}

func takeDamageFromExplotion_two(player *PlayerOne) {
	fmt.Println("Player is taking damage from explosion")
	explosionDmg := 10
	player.health -= explosionDmg
	//return
}

//Access player by pointer
func pointerExampleTwo() {

	player := &PlayerOne{
		health: 100,
	}

	fmt.Printf("Before explotion : %+v\n", player)
	player = nil
	takeDamageFromExplotion_two(player) // pass copy of player
	fmt.Printf("After explotion : %+v\n", player)
}
