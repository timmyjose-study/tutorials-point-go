package main

import ("fmt"
         "strings")

func main() {
  var greeting = "hello, world"
  fmt.Printf("%s\n", greeting)

  for i := 0; i < len(greeting); i++ {
    fmt.Printf("%x ", greeting[i])
  }
  fmt.Println()

  const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
  fmt.Printf("Quoted string is %+q", sampleText) // q to escape unprintable characters, + to escape non-ASCII characters
  fmt.Println()

  // number of bytes
  fmt.Printf("string length = %d\n", len(greeting))

  // almost dynamic in nature - the types of `greetings` 
  // has changed, but Go allows it
  greetings := []string { "Hello", ",", "world" }
  fmt.Printf("%s\n", strings.Join(greetings, " "))
}