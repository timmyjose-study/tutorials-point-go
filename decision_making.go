package main

import "fmt"

func checkGrades(marks int) {
  var grade string = "F"

  switch marks {
    case 90: grade = "A"
    case 80 : grade = "B"
    case 50, 60, 70: grade = "C"
    default: grade = "D"
  }

  switch grade {
    case "A": fmt.Println("Excellent")
    case "B": fmt.Println("Not bad")
    case "C": fmt.Println("Okay")
    case "D": fmt.Println("Bad")
    case "E": fmt.Println("Terrible")
    case "F": fmt.Println("You suck!")
  }
}

func main() {
  var a int = 10

  if (a < 20) {
    fmt.Println("Less than 20")
  } else if (a == 20){
    fmt.Println("Equal to 20")
  } else {
    fmt.Println("Greater than 20")
  }

  fmt.Printf("a is actually %d\n", a)

  checkGrades(90)
}