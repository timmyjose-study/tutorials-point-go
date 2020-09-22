package main

import "fmt"

func main() {
  var a int = 0

  for a <= 20 {
    if a == 11 {
      a++
      continue;
    }
    fmt.Printf("a is %d\n", a)
    a++
  }
}