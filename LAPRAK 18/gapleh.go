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

func cocokUjung(ujung int, kartuBaru Domino) (bool, int) {
	if kartuBaru.sisiA == ujung {
		return true, kartuBaru.sisiB
	}
	if kartuBaru.sisiB == ujung {
		return true, kartuBaru.sisiA
	}
	return false, ujung
}

func mainGapleh(d *Dominoes) ([28]Domino, int) {
	var rangkaian [28]Domino
	var panjangRangkaian int
	var kartuAwal Domino
	var ujungKanan int
	var kartuLewat [28]Domino
	var jumlahLewat int
	var i int
	var kartuBaru Domino
	var cocok bool
	var ujungBaru int
	var masihAda bool

	jumlahLewat = 0
	panjangRangkaian = 0

	if d.jumlahSisa > 0 {
		kartuAwal = ambilKartu(d)
		rangkaian[0] = kartuAwal
		panjangRangkaian = 1
		ujungKanan = kartuAwal.sisiB
	}

	for d.jumlahSisa > 0 {
		kartuBaru = ambilKartu(d)
		cocok, ujungBaru = cocokUjung(ujungKanan, kartuBaru)
		if cocok {
			rangkaian[panjangRangkaian] = kartuBaru
			panjangRangkaian = panjangRangkaian + 1
			ujungKanan = ujungBaru
		} else {
			kartuLewat[jumlahLewat] = kartuBaru
			jumlahLewat = jumlahLewat + 1
		}
	}

	masihAda = true
	for masihAda {
		masihAda = false
		for i = 0; i < jumlahLewat; i++ {
			cocok, ujungBaru = cocokUjung(ujungKanan, kartuLewat[i])
			if cocok {
				rangkaian[panjangRangkaian] = kartuLewat[i]
				panjangRangkaian = panjangRangkaian + 1
				ujungKanan = ujungBaru
				kartuLewat[i] = kartuLewat[jumlahLewat-1]
				jumlahLewat = jumlahLewat - 1
				masihAda = true
				break
			}
		}
	}

	return rangkaian, panjangRangkaian
}

func main() {
	var dominoes Dominoes
	var rangkaian [28]Domino
	var panjang int
	var i int

	fmt.Println("=== Soal 3: Implementasi permainan Gapleh ===")
	kocokKartu(&dominoes)

	fmt.Println("Urutan kartu di tumpukan sebelum bermain:")
	for i = 0; i < 28; i++ {
		fmt.Printf("(%d,%d) ", dominoes.kartu[i].sisiA, dominoes.kartu[i].sisiB)
	}
	fmt.Println()

	rangkaian, panjang = mainGapleh(&dominoes)

	fmt.Println("\nRangkaian kartu Gapleh yang berhasil disambungkan:")
	for i = 0; i < panjang; i++ {
		fmt.Printf("(%d,%d) ", rangkaian[i].sisiA, rangkaian[i].sisiB)
	}
	fmt.Println()
	fmt.Println("Total kartu yang berhasil dirangkai:", panjang, "dari 28 kartu")

	fmt.Println("Rasya Putra Wibowo_109082500132")
}
