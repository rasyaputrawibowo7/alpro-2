package main

import "fmt"

const NMAX = 1000

func selectionSort(A []int, n int) {
	var i, j, idxMin, temp int

	for i = 0; i < n-1; i++ {
		idxMin = i

		for j = i + 1; j < n; j++ {
			if A[j] < A[idxMin] {
				idxMin = j
			}
		}

		temp = A[i]
		A[i] = A[idxMin]
		A[idxMin] = temp
	}
}

func main() {
	var n int
	var m int
	var i, j int
	var rumah [1000]int

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&m)

		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSort(rumah[:], m)

		for j = 0; j < m; j++ {
			fmt.Print(rumah[j])

			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
