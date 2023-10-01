package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Printf("%v\t%T\n", &x, &x)

	s := "James"
	fmt.Printf("%v\t%T\n", &s, &s)

	y := person{
		name: "kavindu",
		age:  23,
	}

	fmt.Printf("%v\t%T\n", &y.age, &y)
}
