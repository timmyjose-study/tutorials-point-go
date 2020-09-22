/* Note that Go uses structural typing instead of nominal typing */

package main

import ("fmt"; "math")

type Circle struct {
  x, y, radius float64
}

// this can be considered a method for the Circle type
func(circle Circle) area() float64 {
  return math.Pi * circle.radius * circle.radius
}

func main() {
  circle := Circle {x : 0, y : -199, radius : 10.0 }
  fmt.Printf("Area of circle = %f\n", circle.area())
}