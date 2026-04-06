package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Deret Fibonacci sampai suku ke-10:")

	for i := 0; i <= 10; i++ {
		fmt.Printf("S_%d = %d\n", i, fibonacci(i))
	}
}