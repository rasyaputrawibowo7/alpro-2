package main

import "fmt"

func main() {
	var suara int
	var jumlahMasuk int
	var jumlahSah int
	var frekuensi [21]int

	var i int
	var ketua int
	var wakil int
	var maks1 int
	var maks2 int

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

	maks1 = -1
	ketua = 0

	for i = 1; i <= 20; i++ {
		if frekuensi[i] > maks1 {
			maks1 = frekuensi[i]
			ketua = i
		}
	}

	maks2 = -1
	wakil = 0

	for i = 1; i <= 20; i++ {
		if i != ketua {
			if frekuensi[i] > maks2 {
				maks2 = frekuensi[i]
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", jumlahMasuk)
	fmt.Println("Suara sah:", jumlahSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
