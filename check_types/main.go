package main

import "fmt"

func checkTypes(val interface{}) {
  switch t := val.(type) {
  case nil:
    fmt.Printf("type of val is %T\n", t)
  case int:
    fmt.Printf("type of val is int\n")

  case float64:
    fmt.Printf("type of val is float64\n")
  case bool, string:
    fmt.Printf("type of val is bool or string\n")
  case func(int) float64:
    fmt.Printf("type of val is a function that takes an int and returns a float64\n")
  default:
    fmt.Println("Unknown type!")
  }
}

func foo(x int) float64 {
  return 0.0
}

func main() {
  checkTypes(100)
  checkTypes(true)
  checkTypes("hello")
  checkTypes(2.133)
  checkTypes(foo)
}