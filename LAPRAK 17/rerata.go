package main

import "fmt"

func main() {
	var angka float64
	var total float64
	var rata float64
	var jumlahData int

	fmt.Scan(&angka)

	for angka != 9999 {
		total += angka
		jumlahData++

		fmt.Scan(&angka)
	}

	if jumlahData != 0 {
		rata = total / float64(jumlahData)
		fmt.Printf("Rerata = %.2f\n", rata)
	} else {
		fmt.Println("Tidak ada data yang dihitung")
	}
}