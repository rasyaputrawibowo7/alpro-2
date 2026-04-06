package main

import "fmt"

func cetakBintang(n int, current int) {
	if current <= n {
		for i := 0; i < current; i++ {
			fmt.Print("*")
		}
		fmt.Println()

		cetakBintang(n, current+1)
	}
}

func main() {
	var n int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	cetakBintang(n, 1)
}