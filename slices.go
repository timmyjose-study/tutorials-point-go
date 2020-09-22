package main

import "fmt"

func printArray(a *[]int) {
  for _, val := range *a {
    fmt.Printf("%d ", val)
  }
  fmt.Println()
}

func main() {
  /* defining a slice:
      
     var numbers []int, and then numbers = []int { 1, 2, 3, 4, 5 }

     or using `make`,

     var numbers = make([]int, 5, 5) // len and cap respectively
   */

   var nums1 []int = []int {1, 2, 3, 4, 5}
   printArray(&nums1)

   var nums2 []int = make([]int, 5, 5)
   for i := 0; i < len(nums2); i++ {
     nums2[i] = i + 1
   }
   printArray(&nums2)

   var nums3 = make([]int, 3, 5)
   fmt.Printf("nums3 has len %d and cap %d with contents %v\n", len(nums3), cap(nums3), nums3)

   var nilSlice []int
   fmt.Printf("nilSlice has len %d and cap %d with contents %v\n", len(nilSlice), cap(nilSlice), nilSlice)

   // append and copy
   var numbers []int
   printSlice(numbers)

   numbers = append(numbers, 0)
   printSlice(numbers)

   numbers = append(numbers, 2, 3, 4)
   printSlice(numbers)

   numbers1 := make([]int, len(numbers), cap(numbers) *2)
   printSlice(numbers1)
   fmt.Println("numebers1 = ", numbers1)
   copy(numbers1, numbers)
   fmt.Println("numebers1 = ", numbers1)
}

func printSlice(slice []int) {
  fmt.Printf("slice = %v, len = %d, cap = %d\n", slice, len(slice), cap(slice))
}