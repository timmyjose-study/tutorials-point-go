package main

import "fmt"

func main() {
  var a int = 10
  var ip *int = &a

  fmt.Printf("Address of a = %x\n", &a)
  fmt.Printf("Address stored in ip = %x\n", ip)
  fmt.Printf("Value stored at the address contained in ip = %d\n", *ip)

  var fp *float32 // default value is nil
  fmt.Printf("The address stored in fp is %x\n", fp)

  // this will induce a segmentation fault
  //fmt.Printf("%f\n", *fp -> this will induce a segmentation fault)
}