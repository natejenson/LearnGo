package main

import (
	"fmt"
	"github.com/natejenson/LearnGo/Structs/shapes"
)

func main() {
	r := shapes.Rectangle{Width: 10, Height: 5}
	c := shapes.Circle{Radius: 12}
	measure(r)
	measure(c)
}

func measure(s shapes.Shape) {
	fmt.Printf("%T %+v has a perimeter of %.3f and area of %.3f\n", s, s, s.Perim(), s.Area())
}
