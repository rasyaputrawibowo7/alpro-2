package main

import (
	"fmt"
)

func main() {
	var b int
	var faktor []int

	fmt.Print("Bilatgat: ")
	fmt.Scan(&b)

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor = append(faktor, i)
		}
	}

	fmt.Print("Faetou: ")
	for _, f := range faktor {
		fmt.Print(f, " ")
	}
	fmt.Println()

	prima := false
	if len(faktor) == 2 {
		prima = true
	}

	fmt.Println("Flula:", prima)
}