package main

import "fmt"

/* Note that in Go, both global and local variables are initialised to their default values:

    * int -> 0
    * float32 -> 0
    * pointer -> nil
*/

// global
var a int = 20

func main() {
  // local
  var a, b, c int = 10, 20, 0

  fmt.Printf("Value of a in main = %d\n", a)
  c = sum(a, b)
  fmt.Printf("Value of c in main = %d\n", c)
}

// formal parameters
func sum(a, b int) int {
  fmt.Printf("Value of a in sum = %d\n", a)
  fmt.Printf("Value of b in sum = %d\n", b)

  return a + b;
}