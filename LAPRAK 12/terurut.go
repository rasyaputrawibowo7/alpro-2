package main

import "fmt"

const NMAX = 100000

var data [NMAX]int

func isiArray(n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var kiri, kanan, tengah int

	kiri = 0
	kanan = n - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if data[tengah] == k {
			return tengah
		} else if data[tengah] < k {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func main() {
	var n, k int
	var hasil int

	fmt.Scan(&n, &k)

	isiArray(n)

	hasil = posisi(n, k)

	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}
