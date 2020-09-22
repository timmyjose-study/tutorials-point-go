package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	l, b int
}

func (r Rectangle) area() float64 {
	return float64(r.l * r.b)
}

type Triangle struct {
	b, h float32
}

func (t Triangle) area() float64 {
	return 0.5 * float64(t.b*t.h)
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0.0, y: 0.0, r: 10.0}
	rectangle := Rectangle{l: 10.0, b: 5.0}
	triangle := Triangle{b: 10.0, h: 5.0}

	fmt.Printf("Area of circle = %f\n", getArea(circle))
	fmt.Printf("Area of rectangle = %f\n", getArea(rectangle))
	fmt.Printf("Area of triangle = %f\n", getArea(triangle))
}
