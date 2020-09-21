package main

import "fmt"

func max(x, y int) int {
  var bigger int

  if x >= y {
    bigger = x
  } else {
    bigger = y
  }

  return bigger
}

func main() {
  var x, y = 10, 20

  fmt.Printf("The bigger value between %d and %d is %d\n", x, y, max(x, y))
}