package main

import "fmt"

func main() {
  var arr [10] int
  var i, j int

  for i = 0; i < len(arr); i++ {
    arr[i] = i + 100
  }

  for j = 0; j < len(arr); j++ {
    fmt.Printf("Element at index %d = %d\n", j, arr[j])
  }
}