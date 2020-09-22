package main

import "fmt"

func main() {
	fmt.Printf("The factorial of %d is %d\n", 10, factorial(10))
	fmt.Printf("The %dth Fibonacci number is %d\n", 11, fib(11))

	for n := 0; n <= 20; n++ {
		fmt.Printf("%d => %d\n", n, fib(n))
	}
}

func factorial(n int) int {
	if n <= 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
