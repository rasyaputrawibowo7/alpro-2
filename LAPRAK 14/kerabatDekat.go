package main

import "fmt"

func selectionSortAsc(A []int, n int) {
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

func selectionSortDesc(A []int, n int) {
	var i, j, idxMax, temp int

	for i = 0; i < n-1; i++ {
		idxMax = i

		for j = i + 1; j < n; j++ {
			if A[j] > A[idxMax] {
				idxMax = j
			}
		}

		temp = A[i]
		A[i] = A[idxMax]
		A[idxMax] = temp
	}
}

func main() {
	var n, m int
	var i, j int
	var rumah [1000]int
	var ganjil [1000]int
	var genap [1000]int
	var jmlGanjil, jmlGenap int

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&m)

		jmlGanjil = 0
		jmlGenap = 0

		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])

			if rumah[j]%2 == 1 {
				ganjil[jmlGanjil] = rumah[j]
				jmlGanjil++
			} else {
				genap[jmlGenap] = rumah[j]
				jmlGenap++
			}
		}

		selectionSortAsc(ganjil[:], jmlGanjil)
		selectionSortDesc(genap[:], jmlGenap)

		for j = 0; j < jmlGanjil; j++ {
			fmt.Print(ganjil[j], " ")
		}

		for j = 0; j < jmlGenap; j++ {
			fmt.Print(genap[j])

			if j < jmlGenap-1 {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}
