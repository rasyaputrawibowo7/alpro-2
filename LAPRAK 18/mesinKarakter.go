//Rasya Putra Wibowo_109082500132
package main

import "fmt"

type MesinKarakter struct {
	teks   string
	posisi int
}

func start(m *MesinKarakter, teksBaru string) {
	m.teks = teksBaru
	m.posisi = 0
}

func maju(m *MesinKarakter) {
	m.posisi = m.posisi + 1
}

func eop(m *MesinKarakter) bool {
	if m.posisi >= len(m.teks) {
		return true
	}
	if m.teks[m.posisi] == '.' {
		return true
	}
	return false
}

func cc(m *MesinKarakter) byte {
	if m.posisi >= len(m.teks) {
		return '.'
	}
	return m.teks[m.posisi]
}

func hitungKarakter(m *MesinKarakter, teks string) int {
	var jumlah int
	jumlah = 0
	start(m, teks)
	for eop(m) == false {
		jumlah = jumlah + 1
		maju(m)
	}
	return jumlah
}

func hitungHurufA(m *MesinKarakter, teks string) int {
	var jumlahA int
	jumlahA = 0
	start(m, teks)
	for eop(m) == false {
		if cc(m) == 'A' {
			jumlahA = jumlahA + 1
		}
		maju(m)
	}
	return jumlahA
}

func frekuensiHurufA(m *MesinKarakter, teks string) int {
	var totalKarakter int
	var totalA int
	totalKarakter = hitungKarakter(m, teks)
	totalA = hitungHurufA(m, teks)
	if totalKarakter == 0 {
		return 0
	}
	return (totalA * 100) / totalKarakter
}

func hitungKataLE(m *MesinKarakter, teks string) int {
	var jumlahLE int
	var karakterSebelumnya byte
	var adaKarakterSebelumnya bool

	jumlahLE = 0
	adaKarakterSebelumnya = false
	start(m, teks)
	for eop(m) == false {
		if adaKarakterSebelumnya {
			if karakterSebelumnya == 'L' && cc(m) == 'E' {
				jumlahLE = jumlahLE + 1
			}
		}
		karakterSebelumnya = cc(m)
		adaKarakterSebelumnya = true
		maju(m)
	}
	return jumlahLE
}

func main() {
	var mesin MesinKarakter
	var teksUji string

	fmt.Println("=== Soal 4: Mesin abstrak karakter ===")
	fmt.Println("Masukkan untaian karakter (diakhiri tanda titik '.'):")
	fmt.Scanln(&teksUji)

	fmt.Println("\n[4b-i] Jumlah karakter yang terbaca (tanpa tanda titik):")
	fmt.Println(hitungKarakter(&mesin, teksUji))

	fmt.Println("\n[4b-ii] Jumlah huruf 'A' yang terbaca:")
	fmt.Println(hitungHurufA(&mesin, teksUji))

	fmt.Println("\n[4b-iii] Frekuensi huruf 'A' terhadap seluruh karakter (dalam persen):")
	fmt.Println(frekuensiHurufA(&mesin, teksUji))

	fmt.Println("\n[4b-iv] Jumlah kata 'LE' (pasangan berurutan huruf L dan E) yang terbaca:")
	fmt.Println(hitungKataLE(&mesin, teksUji))

	fmt.Println("Rasya Putra Wibowo_109082500132")
}
