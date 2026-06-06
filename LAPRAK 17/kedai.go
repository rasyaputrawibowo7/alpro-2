package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var jumlahTopping int
	var nomor int
	var posisiX, posisiY float64
	var toppingDiPizza int
	var estimasiPi float64

	fmt.Print("Banyak Topping: ")
	fmt.Scan(&jumlahTopping)

	toppingDiPizza = 0

	for nomor = 1; nomor <= jumlahTopping; nomor++ {
		posisiX = rand.Float64()
		posisiY = rand.Float64()

		if (posisiX-0.5)*(posisiX-0.5)+(posisiY-0.5)*(posisiY-0.5) <= 0.25 {
			toppingDiPizza++
		}
	}

	estimasiPi = 4.0 * float64(toppingDiPizza) / float64(jumlahTopping)

	fmt.Println("Topping pada Pizza:", toppingDiPizza)
	fmt.Printf("PI : %.10f\n", estimasiPi)
}
