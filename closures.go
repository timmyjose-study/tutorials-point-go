package main

import "fmt"

func getSequence() func() int {
  i := 0

  return func() int {
    i += 1
    return i
  }
}

func addN(n int) func(x int) int {
  return func(x int) int {
    return x + n
  }
}

func main() {
  nextNumber := getSequence()

  for i := 0; i < 10; i++ {
    fmt.Println(nextNumber())
  }

  nextNumber = getSequence()
  for j := 0; j < 5; j++ {
    fmt.Println(nextNumber())
  }

  fmt.Println(addN(100)(1))
  fmt.Println(addN(12)(1))
  fmt.Println(addN(0)(1))
  fmt.Println(addN(9999)(1))
  fmt.Println(addN(-10)(1))
}