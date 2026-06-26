//Rasya Putra Wibowo_109082500132
package main

import "fmt"

type Domino struct {
	sisiA int
	sisiB int
	nilai int
	balak bool
}

type Dominoes struct {
	kartu      [28]Domino
	jumlahSisa int
}

func kocokKartu(d *Dominoes) {
	var i int
	var j int
	var indeks int
	var seed int

	indeks = 0
	for i = 0; i <= 6; i++ {
		for j = i; j <= 6; j++ {
			d.kartu[indeks].sisiA = i
			d.kartu[indeks].sisiB = j
			d.kartu[indeks].nilai = i + j
			if i == j {
				d.kartu[indeks].balak = true
			} else {
				d.kartu[indeks].balak = false
			}
			indeks = indeks + 1
		}
	}
	d.jumlahSisa = 28

	fmt.Println("Masukkan angka seed (bilangan bulat) untuk mengacak kartu:")
	fmt.Scan(&seed)

	for i = 27; i >= 1; i-- {
		seed = (seed*1103515245 + 12345) % 2147483648
		if seed < 0 {
			seed = seed * -1
		}
		j = seed % (i + 1)

		var temp Domino
		temp = d.kartu[i]
		d.kartu[i] = d.kartu[j]
		d.kartu[j] = temp
	}
}

func ambilKartu(d *Dominoes) Domino {
	var hasil Domino
	if d.jumlahSisa <= 0 {
		hasil.sisiA = -1
		hasil.sisiB = -1
		hasil.nilai = -1
		hasil.balak = false
		return hasil
	}
	hasil = d.kartu[d.jumlahSisa-1]
	d.jumlahSisa = d.jumlahSisa - 1
	return hasil
}

func gambarKartu(kartu Domino, suit int) int {
	if kartu.sisiA == suit {
		return kartu.sisiB
	}
	if kartu.sisiB == suit {
		return kartu.sisiA
	}
	return -1
}

func nilaiKartu(kartu Domino) int {
	return kartu.nilai
}

func main() {
	var dominoes Dominoes
	var i int
	var jumlahAmbil int
	var kartuDiambil [28]Domino
	var suitInput int

	fmt.Println("=== Soal 1: Implementasi mesin abstrak kartu domino ===")

	kocokKartu(&dominoes)

	fmt.Println("\n[1b] Urutan 28 kartu setelah kocokKartu:")
	for i = 0; i < 28; i++ {
		fmt.Printf("(%d,%d) ", dominoes.kartu[i].sisiA, dominoes.kartu[i].sisiB)
	}
	fmt.Println()
	fmt.Println("Jumlah kartu sisa:", dominoes.jumlahSisa)

	fmt.Println("\n[1c] Masukkan jumlah kartu yang ingin diambil dari tumpukan:")
	fmt.Scan(&jumlahAmbil)

	for i = 0; i < jumlahAmbil; i++ {
		kartuDiambil[i] = ambilKartu(&dominoes)
		fmt.Printf("Kartu ke-%d diambil: (%d,%d), nilai=%d, balak=%t\n",
			i+1, kartuDiambil[i].sisiA, kartuDiambil[i].sisiB, kartuDiambil[i].nilai, kartuDiambil[i].balak)
	}
	fmt.Println("Jumlah kartu sisa setelah pengambilan:", dominoes.jumlahSisa)

	fmt.Println("\n[1d] Masukkan nilai suit (gambar) yang ingin dicari pada kartu pertama yang diambil:")
	fmt.Scan(&suitInput)
	fmt.Printf("gambarKartu pada (%d,%d) dengan suit=%d menghasilkan: %d\n",
		kartuDiambil[0].sisiA, kartuDiambil[0].sisiB, suitInput, gambarKartu(kartuDiambil[0], suitInput))

	fmt.Println("\n[1e] Nilai dari setiap kartu yang diambil:")
	for i = 0; i < jumlahAmbil; i++ {
		fmt.Printf("nilaiKartu(%d,%d) = %d\n", kartuDiambil[i].sisiA, kartuDiambil[i].sisiB, nilaiKartu(kartuDiambil[i]))
	}

	fmt.Println("Rasya Putra Wibowo_109082500132")
}
