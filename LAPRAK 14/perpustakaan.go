package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	var i int

	fmt.Scan(n)

	for i = 0; i < *n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var iMaks, i int

	iMaks = 0

	for i = 1; i < n; i++ {
		if pustaka[i].rating > pustaka[iMaks].rating {
			iMaks = i
		}
	}

	fmt.Println(
		pustaka[iMaks].judul,
		pustaka[iMaks].penulis,
		pustaka[iMaks].penerbit,
		pustaka[iMaks].tahun,
	)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var pass, i int
	var temp Buku

	for pass = 1; pass < n; pass++ {
		temp = pustaka[pass]
		i = pass - 1

		for i >= 0 && pustaka[i].rating < temp.rating {
			pustaka[i+1] = pustaka[i]
			i--
		}

		pustaka[i+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i, batas int

	if n < 5 {
		batas = n
	} else {
		batas = 5
	}

	for i = 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	var kiri, kanan, tengah int

	kiri = 0
	kanan = n - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if pustaka[tengah].rating == r {
			fmt.Println(
				pustaka[tengah].judul,
				pustaka[tengah].penulis,
				pustaka[tengah].penerbit,
				pustaka[tengah].tahun,
				pustaka[tengah].eksemplar,
				pustaka[tengah].rating,
			)
			return
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	fmt.Println("Tidak ada buku dengan rating seperti itu")
}

func main() {
	var pustaka DaftarBuku
	var n int
	var ratingCari int

	DaftarkanBuku(&pustaka, &n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	fmt.Scan(&ratingCari)

	CariBuku(pustaka, n, ratingCari)
}
