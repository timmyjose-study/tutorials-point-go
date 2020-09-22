package main

import "fmt"

func invoke(x int, f *func(int) int) {
  fmt.Printf("%d\n", (*f)(x))
}

func invokeAgain(x int, f func(int) int) {
  fmt.Printf("%d\n", f(x))
}

func main() {
  var squarer = func(x int) int {
    return x * x
  }

  invoke(10, &squarer)
  invokeAgain(10, squarer)
}
