package main

import "fmt"

func main() {
	var K int
	var hasil float64 = 1

	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	for k := 0; k <= K; k++ {
		pembilang := float64((4*k + 2) * (4*k + 2))
		penyebut := float64((4*k + 1) * (4*k + 3))
		hasil *= pembilang / penyebut
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}