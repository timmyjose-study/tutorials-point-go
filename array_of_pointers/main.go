package main

import "fmt"

const MAX int = 3

func main() {
  a := []int {10, 20, 999 }

  var ptr [MAX]*int

  for i := 0; i < len(a); i++ {
    ptr[i] = &a[i]
  }

  for j := 0; j < len(ptr); j++ {
    fmt.Printf("%d\n", *ptr[j])
  }
}