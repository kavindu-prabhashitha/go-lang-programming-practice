package main

import "fmt"

// 8 bytes => pointer (64 bit machine)
// 4 bytes => pointer (32 bit machine)

type PlayerOne struct {
	health int
}

func takeDamageFromExplotion_one(player *PlayerOne) {
	fmt.Println("Player is taking damage from explosion")
	explosionDmg := 10
	player.health -= explosionDmg
	//return
}

//Access player by pointer
func pointerExampleOne() {

	player := PlayerOne{
		health: 100,
	}

	fmt.Printf("Before explotion : %+v\n", player)
	takeDamageFromExplotion_one(&player) // pass copy of player
	fmt.Printf("After explotion : %+v\n", player)
}
