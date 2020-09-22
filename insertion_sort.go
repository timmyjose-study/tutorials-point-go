package main

import "fmt"

func insertionSort(arr []int) {
  for i := 1; i < len(arr); i++ {
    k := arr[i]
    j := i - 1

    for ; j >= 0 && arr[j] > k; j-- {
      arr[j+1] = arr[j]
    }
    arr[j+1] = k
  }
}

func print(arr []int) {
  for _, val := range arr {
    fmt.Printf("%d ", val)
  }
  fmt.Println()
}

func main() {
  var numbers = []int { 1, 0, 2, 3, 4, 3, 5, 99, -100, 11 }
  print(numbers)

  insertionSort(numbers)
  print(numbers)
}