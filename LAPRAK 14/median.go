package main

import "fmt"

const NMAX = 1000000

func insertionSort(A []int, n int) {
	var i, pass, temp int

	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass - 1

		for i >= 0 && A[i] > temp {
			A[i+1] = A[i]
			i--
		}

		A[i+1] = temp
	}
}

func median(A []int, n int) int {
	if n%2 == 1 {
		return A[n/2]
	}

	return (A[(n/2)-1] + A[n/2]) / 2
}

func main() {
	var data [NMAX]int
	var n int
	var x int
	var hasil int

	n = 0

	for {
		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x == 0 {
			insertionSort(data[:], n)
			hasil = median(data[:], n)
			fmt.Println(hasil)
		} else {
			data[n] = x
			n++
		}
	}
}
