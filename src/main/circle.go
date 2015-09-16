package main

import (
	//"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func area(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

//func main() {
//var circle = Circle{1, 2, 3}
//fmt.Printf("area = %f", circle.area())
//fmt.Printf("area = %f", area(&circle))
//}
