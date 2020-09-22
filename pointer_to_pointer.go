package main

import "fmt"

func main() {
  var a int = 100
  var pi *int = &a
  var ppi **int = &pi

  fmt.Printf("a = %d, pi = %d, ppi = %d\n", a, *pi, **ppi)
}