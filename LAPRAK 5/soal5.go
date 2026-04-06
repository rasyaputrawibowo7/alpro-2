package main

import "fmt"

func printOdd(current, n int) {
	if current > n {
		return
	}

	if current%2 != 0 {
		fmt.Print(current, " ")
	}

	printOdd(current+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	printOdd(1, n)
}