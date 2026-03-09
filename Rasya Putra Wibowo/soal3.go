package main

import (
	"fmt"
)

func main() {
	var r int
	var volume, luas float64
	const pi = 3.1415926535

	fmt.Print("Jari-jari: ")
	fmt.Scan(&r)

	rf := float64(r)

	volume = (4.0 / 3.0) * pi * rf * rf * rf
	luas = 4 * pi * rf * rf

	fmt.Printf("Volume bola = %.4f\n", volume)
	fmt.Printf("Luas kulit bola = %.4f\n", luas)
}