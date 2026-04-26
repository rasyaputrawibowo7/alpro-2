package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var ikan [1000]float64

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var wadah [1000]float64
	index := 0

	for i := 0; i < x; i += y {
		total := 0.0
		for j := i; j < i+y && j < x; j++ {
			total += ikan[j]
		}
		wadah[index] = total
		index++
	}

	for i := 0; i < index; i++ {
		fmt.Printf("%.2f ", wadah[i])
	}
	fmt.Println()

	sum := 0.0
	for i := 0; i < index; i++ {
		sum += wadah[i]
	}
	rata := sum / float64(index)

	fmt.Printf("%.2f\n", rata)
}