package main

import "fmt"

func main() {
	z := 42
	var x = 40
	const y = 41

	fmt.Printf("Value of z is %v and the type of Z is %T \n", z, z)
	fmt.Printf("Value of x is %v and the type of x is %T \n", x, x)
	fmt.Printf("Value of y is %v and the type of y is %T \n", y, y)
}
