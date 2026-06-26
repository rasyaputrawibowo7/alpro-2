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

func galiKartu(d *Dominoes, kartuAcuan Domino) Domino {
	var kartuDigali Domino
	var ditemukan bool

	ditemukan = false
	for d.jumlahSisa > 0 && ditemukan == false {
		kartuDigali = ambilKartu(d)
		if kartuDigali.sisiA == kartuAcuan.sisiA || kartuDigali.sisiA == kartuAcuan.sisiB ||
			kartuDigali.sisiB == kartuAcuan.sisiA || kartuDigali.sisiB == kartuAcuan.sisiB {
			ditemukan = true
		}
	}

	if ditemukan == false {
		kartuDigali.sisiA = -1
		kartuDigali.sisiB = -1
		kartuDigali.nilai = -1
		kartuDigali.balak = false
	}
	return kartuDigali
}

func sepasangKartu(kartu1 Domino, kartu2 Domino) bool {
	var total int
	total = kartu1.nilai + kartu2.nilai
	if total == 12 {
		return true
	}
	return false
}

func main() {
	var dominoes Dominoes
	var kartuAcuan Domino
	var kartuGalian Domino
	var sisiAInput int
	var sisiBInput int

	fmt.Println("=== Soal 2: Realisasi aksi galiKartu & sepasangKartu ===")

	kocokKartu(&dominoes)

	fmt.Println("\n[2a] Masukkan sisiA kartu acuan (0-6):")
	fmt.Scan(&sisiAInput)
	fmt.Println("Masukkan sisiB kartu acuan (0-6):")
	fmt.Scan(&sisiBInput)

	kartuAcuan.sisiA = sisiAInput
	kartuAcuan.sisiB = sisiBInput
	kartuAcuan.nilai = sisiAInput + sisiBInput
	if sisiAInput == sisiBInput {
		kartuAcuan.balak = true
	} else {
		kartuAcuan.balak = false
	}

	fmt.Printf("Kartu acuan: (%d,%d)\n", kartuAcuan.sisiA, kartuAcuan.sisiB)
	fmt.Println("Sisa kartu di tumpukan sebelum menggali:", dominoes.jumlahSisa)

	kartuGalian = galiKartu(&dominoes, kartuAcuan)
	fmt.Printf("Kartu hasil galiKartu: (%d,%d)\n", kartuGalian.sisiA, kartuGalian.sisiB)
	fmt.Println("Sisa kartu di tumpukan setelah menggali:", dominoes.jumlahSisa)

	fmt.Println("\n[2b] Uji sepasangKartu pada kartu acuan dan kartu hasil galian:")
	fmt.Printf("sepasangKartu((%d,%d), (%d,%d)) = %t (nilai1=%d + nilai2=%d = %d)\n",
		kartuAcuan.sisiA, kartuAcuan.sisiB, kartuGalian.sisiA, kartuGalian.sisiB,
		sepasangKartu(kartuAcuan, kartuGalian),
		kartuAcuan.nilai, kartuGalian.nilai, kartuAcuan.nilai+kartuGalian.nilai)

	fmt.Println("Rasya Putra Wibowo_109082500132")
}
