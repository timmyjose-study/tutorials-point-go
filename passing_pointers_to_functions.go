package main

import "fmt"

func main() {
  var a, b = 10, 20

  fmt.Printf("Before swap, a = %d, b = %d\n", a, b)
  swap(&a, &b)
  fmt.Printf("After swap, a = %d, b = %d\n", a, b)
}

func swap(x, y *int) {
  t := *x
  *x = *y
  *y = t
}