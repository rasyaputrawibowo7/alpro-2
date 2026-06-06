package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var jumlahTetes int
	var urutan int
	var koordinatX, koordinatY float64
	var wilayah1, wilayah2, wilayah3, wilayah4 int

	fmt.Print("Masukkan banyak tetesan hujan: ")
	fmt.Scan(&jumlahTetes)

	wilayah1 = 0
	wilayah2 = 0
	wilayah3 = 0
	wilayah4 = 0

	for urutan = 1; urutan <= jumlahTetes; urutan++ {
		koordinatX = rand.Float64()
		koordinatY = rand.Float64()

		if koordinatX < 0.5 && koordinatY < 0.5 {
			wilayah1++
		} else if koordinatX >= 0.5 && koordinatY < 0.5 {
			wilayah2++
		} else if koordinatX >= 0.5 && koordinatY >= 0.5 {
			wilayah3++
		} else {
			wilayah4++
		}
	}

	fmt.Printf("Curah hujan wilayah 1: %.4f milimeter\n", float64(wilayah1)*0.0001)
	fmt.Printf("Curah hujan wilayah 2: %.4f milimeter\n", float64(wilayah2)*0.0001)
	fmt.Printf("Curah hujan wilayah 3: %.4f milimeter\n", float64(wilayah3)*0.0001)
	fmt.Printf("Curah hujan wilayah 4: %.4f milimeter\n", float64(wilayah4)*0.0001)
}
