/* Golang has goto! */

package main

import "fmt"

func main() {
  var a int = 0

  LOOP: for a <= 20 {
    if a == 11 {
      a++
      goto LOOP
    }
    fmt.Printf("a is %d\n", a)
    a++
  }
}