package main

import "fmt"

func main() {
  var b int = 5
  var a int
  numbers := [6]int {1, 2, 3, 5}

  // initial value of a is 0
  fmt.Printf("a is initially %d\n", a)

  // this a is local
  for a := 0; a < 10; a++ {
    fmt.Printf("value of a: %d\n", a)
  }

  for a < b {
    a++
    fmt.Printf("New value of a: %d\n", a)
  }

  for idx, num := range numbers {
    fmt.Printf("Number at index %d = %d\n", idx, num)
  }
}