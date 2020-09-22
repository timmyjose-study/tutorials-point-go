package main

import "fmt"

func main() {
	numbers := []int{11, 12, 13, 14, 15}

	// with arrays/slices
	for _, num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// with maps
	phoneBook := map[string]int{"Bob": 100, "Dave": 200, "Susan": 300}

	for name, phone := range phoneBook {
		fmt.Printf("%s => %d\n", name, phone)
	}

	// with strings
	message := "hello, world"
	for _, c := range message {
		fmt.Printf("%c ", c)
	}
	fmt.Println()
}
