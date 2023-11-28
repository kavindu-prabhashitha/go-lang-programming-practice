package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p Point) scale(s float64) {
	p.x = p.x * s
	p.y = p.y * s
}

func (p *Point) size() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main() {
	p := &Point{3, 4}
	fmt.Printf("before scaling: = %+v, size: %v\n", p, p.size())
	var sc float64 = 5
	p.scale(sc)
	fmt.Printf("after scaling: %+v, size: %v\n", p, p.size())
}

// Methods with pointer receivers can modify the value to which the receiver points (as scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
