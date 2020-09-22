package main

import "fmt"

func getAverage(arr []int) float32 {
  var sum float32

  for _, val := range arr {
    sum += float32(val)
  }

  return sum / float32 (len (arr))
}

func main() {
  var numbers = []int { 1, 12, 3, 14, 15, 6, 17 }
  fmt.Printf("The average of the array is %f\n", getAverage(numbers))
}