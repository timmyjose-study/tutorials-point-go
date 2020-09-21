package main

import ("fmt"; "math")

func main() {
  // lambda function
  getSquareRoot := func(x float64) float64 {
    return math.Sqrt(x)
  }

  fmt.Printf("The square root of %f is %f\n", 19, getSquareRoot(19))
}