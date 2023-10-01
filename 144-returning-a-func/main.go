package main

import "fmt"

func main() {

	x := foo()
	fmt.Println(x)

	y := bar()
	fmt.Println(y())

	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", y)

	z := maa(23)
	fmt.Println(z())
}

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 43
	}
}

func maa(s int) func() int {

	zz := func() {
		fmt.Printf("Hello from inside function")
	}

	zz()

	return func() int {
		return 43 + s
	}
}
