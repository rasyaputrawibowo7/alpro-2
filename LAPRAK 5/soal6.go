package main

import "fmt"

func power(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * power(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y)

	result := power(x, y)
	fmt.Println("Hasil:", result)
}