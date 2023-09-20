package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("this is the first statment to run")
	fmt.Println("this is the second statement to run")
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement
	fmt.Printf("x = %v \n y = %v \n", x, y)

	// CONDITIONAL
	// if statement
	// switch statements

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is Greater than or equal x which is %v \n", z, x)
	} else {
		fmt.Printf("z is %v and that is Less than  which is %v \n", z, x)
	}
}
