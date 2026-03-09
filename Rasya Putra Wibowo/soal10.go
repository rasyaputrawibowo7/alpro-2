package main

import "fmt"

func main() {
	var berat int
	var kg, sisa int
	var biayaKg, biayaGram, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	sisa = berat % 1000

	biayaKg = kg * 10000

	if sisa >= 500 {
		biayaGram = sisa * 5
	} else {
		biayaGram = sisa * 15
	}

	if kg > 10 {
		biayaGram = 0
	}

	total = biayaKg + biayaGram

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gram")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaGram)
	fmt.Println("Total biaya: Rp.", total)
}