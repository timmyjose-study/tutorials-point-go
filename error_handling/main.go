package main

import (
	"errors"
	"fmt"
)

func safeDivision(m, n int) (int, error) {
	if n == 0 {
		return 0, errors.New("Some people just love to see the world burn")
	}
	return m / n, nil
}

func main() {
	validRes, err := safeDivision(10, 2)

	if err != nil {
		fmt.Printf("Error while attempting division: %s\n", err)
	} else {
		fmt.Printf("%d\n", validRes)
	}

	defaultRes, err := safeDivision(10, 0)
	if err != nil {
		fmt.Printf("Error while attempting division: %s\n", err)
	} else {
		fmt.Printf("%d\n", defaultRes)
	}
}
