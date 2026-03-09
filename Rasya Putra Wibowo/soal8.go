package main

import (
	"fmt"
)

func main() {
	var kiri, kanan float64
	var selisih float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)

		if kiri < 0 || kanan < 0 || kiri+kanan > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		if kiri > kanan {
			selisih = kiri - kanan
		} else {
			selisih = kanan - kiri
		}

		fmt.Print("Sepeda motor pak Andi akan oleng: ")
		fmt.Println(selisih >= 9)
	}
}