package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	// Input array
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	// a. Tampilkan semua isi array
	fmt.Println("\nSemua isi array:")
	fmt.Println(arr)

	// b. Indeks ganjil
	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}

	// c. Indeks genap (index 0 termasuk genap)
	fmt.Println("\n\nElemen dengan indeks genap:")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	// d. Indeks kelipatan x
	var x int
	fmt.Print("\n\nMasukkan nilai x (kelipatan indeks): ")
	fmt.Scan(&x)

	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	// e. Hapus elemen pada indeks tertentu
	var idx int
	fmt.Print("\n\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)

	arr = append(arr[:idx], arr[idx+1:]...)

	fmt.Println("Array setelah penghapusan:")
	fmt.Println(arr)

	// f. Rata-rata
	sum := 0
	for _, v := range arr {
		sum += v
	}
	mean := float64(sum) / float64(len(arr))
	fmt.Println("\nRata-rata:", mean)

	// g. Standar deviasi
	var variance float64
	for _, v := range arr {
		variance += math.Pow(float64(v)-mean, 2)
	}
	variance /= float64(len(arr))
	stdDev := math.Sqrt(variance)
	fmt.Println("Standar deviasi:", stdDev)

	// h. Frekuensi angka tertentu
	var target, count int
	fmt.Print("\nMasukkan angka yang ingin dicari frekuensinya: ")
	fmt.Scan(&target)

	for _, v := range arr {
		if v == target {
			count++
		}
	}
	fmt.Println("Frekuensi:", count)
}