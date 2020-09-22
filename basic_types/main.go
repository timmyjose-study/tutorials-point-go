package main

import "fmt"

func main() {
  var x float64
  x = 20.0
  fmt.Println(x)
  fmt.Printf("x is of type %T\n", x)

  var y uint32
  y = 1888
  fmt.Printf("%d is of type %T\n", y, y)

  // type inference
  z := "hello"
  fmt.Printf("\"%s\" has the type %T\n", z, z)

  var a, b, c = 1, 3.0, "hi"
  fmt.Printf("%d, %f, %s\n", a, b, c)
}