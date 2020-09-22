package main

import "fmt"

func main() {
  numbers := []int { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
  printSlice(numbers)
  fmt.Println("numbers = ", numbers)
  fmt.Println("numbers[1:4] = ", numbers[1:4])
  fmt.Println("numbers[:3] = ", numbers[:3])
  fmt.Println("numbers[4:] = ", numbers[4:])

}

func printSlice(a []int) {
  fmt.Printf("slice = %v, len = %d, cap = %d\n", a, len(a), cap(a))
}