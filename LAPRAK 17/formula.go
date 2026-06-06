package main

import "fmt"

func main() {
	var nSuku int
	var k int
	var nilai float64
	var hasil float64

	fmt.Print("Masukkan jumlah suku: ")
	fmt.Scan(&nSuku)

	hasil = 0

	for k = 1; k <= nSuku; k++ {
		if k%2 != 0 {
			nilai = 1.0 / float64(2*k-1)
		} else {
			nilai = -1.0 / float64(2*k-1)
		}

		hasil += nilai
	}

	fmt.Printf("Nilai PI: %.7f\n", 4*hasil)

	k = 1
	hasil = 0

	for {
		if k%2 != 0 {
			nilai = 1.0 / float64(2*k-1)
		} else {
			nilai = -1.0 / float64(2*k-1)
		}

		hasil += nilai

		if (1.0 / float64(2*k+1)) <= 0.00001 {
			break
		}

		k++
	}

	fmt.Printf("Nilai PI: %.10f\n", 4*hasil)
	fmt.Println("Iterasi berhenti pada suku ke-", k)
}
