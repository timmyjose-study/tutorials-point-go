package main

import "fmt"

func main() {
  var numbers = []float32 { 1.0, 2.0, 3.0, 4.0, 5.0 }

  for idx, val := range numbers {
    fmt.Printf("Value at index %d = %f\n", idx, val)
  }
}