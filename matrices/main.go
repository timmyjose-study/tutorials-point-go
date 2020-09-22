package main

import "fmt"

func main() {
  var m = [3][4] int {
    {0, 1, 2, 3},
    {4, 5, 6, 7},
    {8, 9, 10, 11},
  }

  for _, row := range m {
    for _, val := range row {
      fmt.Printf("%d ", val)
    }
    fmt.Println()
  }
}