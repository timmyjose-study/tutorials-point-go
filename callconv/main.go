/* Golang has Call-by-value by default, can also use Call-by-reference */

package main

import "fmt"

func BadSwap(x, y int) {
  temp := x
  x = y
  y = temp
}

func GoodSwap(px, py *int) {
  temp := *px
  *px = *py
  *py = temp
}

func main() {
  var x, y int = 10, 20

  fmt.Printf("Originally, x = %d, y = %d\n", x, y)
  BadSwap(x, y)
  fmt.Printf("After BadSwap, x = %d, y = %d\n", x, y)
  GoodSwap(&x, &y)
  fmt.Printf("After GoodSwap, x = %d, y = %d\n", x, y)
}