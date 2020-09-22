package main

import "fmt"

// global variables
var g int

func main() {
  a, b := 10, 20
  g = a + b

  fmt.Printf("The sum of %d and %d is %d\n", a, b, g)

  // Go allows shadowing
  {
    var g = 100;
    fmt.Printf("g inside one level = %d\n", g)

    {
      var g = 200;
      fmt.Printf("g inside two levels = %d\n", g)
    }

    {
      var g = 300;
      fmt.Printf("g inside three levels = %d\n", g)
    }
  }
}