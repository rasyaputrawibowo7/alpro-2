package main

import "fmt"

const NMAX = 1000

func insertionSort(A []int, n int) {
	var pass, i, temp int

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

func main() {
	var data [NMAX]int
	var n int
	var x int
	var jarak int
	var tetap bool
	var i int

	n = 0

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		data[n] = x
		n++
	}

	insertionSort(data[:], n)

	for i = 0; i < n; i++ {
		fmt.Print(data[i])

		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	tetap = true

	if n > 1 {
		jarak = data[1] - data[0]

		for i = 2; i < n; i++ {
			if data[i]-data[i-1] != jarak {
				tetap = false
			}
		}

		if tetap {
			fmt.Println("Data berjarak", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	} else {
		fmt.Println("Data berjarak 0")
	}
}
