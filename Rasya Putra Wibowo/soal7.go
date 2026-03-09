package main

import (
	"fmt"
)

func main() {
	var bunga string
	var pita string
	var jumlah int
	i := 1

	for {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scanln(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita = pita + bunga + " - "
		jumlah++
		i++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}