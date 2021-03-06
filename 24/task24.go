package main24

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p1 *Point) Distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x,2) + math.Pow(p1.y-p2.y,2))
}
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}
func Main24() {
	p1:=NewPoint(3,4)
	p2:=NewPoint(6,8)
	fmt.Println(p1.Distance(p2))
}