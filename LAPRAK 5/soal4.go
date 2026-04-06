package main

import "fmt"

func printSequence(n int, current int) {
	if current == 0 {
		return
	}

	fmt.Print(current, " ")

	printSequence(n, current-1)

	if current != n {
		fmt.Print(current, " ")
	}
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)

	printSequence(n, n)
}