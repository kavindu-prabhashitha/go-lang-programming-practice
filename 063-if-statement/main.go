package main

import "fmt"

func main() {
	fmt.Println("this is the first statment to run")
	fmt.Println("this is the second statement to run")
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement
	fmt.Printf("x = %v \n y = %v \n", x, y)

	// CONDITIONAL
	// if statement
	// switch statements

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("equal to , or greater than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal to the meaning of life")
	} else {
		fmt.Println("equal to , or greater than the meaning of life")
	}
}
