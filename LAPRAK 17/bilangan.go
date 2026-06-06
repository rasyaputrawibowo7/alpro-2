package main

import "fmt"

func main() {
	var kataCari, kata string
	var banyakData, indeks int
	var ada bool
	var letakPertama int
	var totalKemunculan int

	fmt.Print("Masukkan kata yang dicari: ")
	fmt.Scan(&kataCari)

	fmt.Print("Masukkan banyak data: ")
	fmt.Scan(&banyakData)

	ada = false
	letakPertama = -1
	totalKemunculan = 0

	for indeks = 1; indeks <= banyakData; indeks++ {
		fmt.Print("Masukan data ke-", indeks, ": ")
		fmt.Scan(&kata)

		if kata == kataCari {
			totalKemunculan++
			ada = true

			if letakPertama == -1 {
				letakPertama = indeks
			}
		}
	}

	fmt.Println()
	fmt.Println("Hasil Pemeriksaan:")

	if ada {
		fmt.Println("a. Kata yang dicari ditemukan dalam data.")
	} else {
		fmt.Println("a. Kata yang dicari tidak ditemukan dalam data.")
	}

	if letakPertama != -1 {
		fmt.Println("b. Kemunculan pertama berada pada posisi", letakPertama)
	} else {
		fmt.Println("b. Kata yang dicari tidak ditemukan.")
	}

	fmt.Println("c. Jumlah kemunculan kata =", totalKemunculan)

	if totalKemunculan >= 2 {
		fmt.Println("d. Terdapat minimal dua kemunculan kata tersebut.")
	} else {
		fmt.Println("d. Kemunculan kata kurang dari dua kali.")
	}
}
