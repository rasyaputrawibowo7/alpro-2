package main

import "fmt"

func main() {
	var suara int
	var jumlahMasuk int
	var jumlahSah int
	var frekuensi [21]int
	var i int

	jumlahMasuk = 0
	jumlahSah = 0
	
	for {
		fmt.Scan(&suara)

		if suara == 0 {
			break
		}

		jumlahMasuk++

		if suara >= 1 && suara <= 20 {
			jumlahSah++
			frekuensi[suara]++
		}
	}

	fmt.Println("Suara masuk:", jumlahMasuk)
	fmt.Println("Suara sah:", jumlahSah)

	for i = 1; i <= 20; i++ {
		if frekuensi[i] > 0 {
			fmt.Println(i, ":", frekuensi[i])
		}
	}
}