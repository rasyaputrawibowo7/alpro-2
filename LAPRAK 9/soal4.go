package main

import (
	"fmt"
)

const NMAX = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	fmt.Print("Teks: ")
	var input string
	fmt.Scan(&input)

	*n = len(input)
	for i := 0; i < *n; i++ {
		t[i] = rune(input[i])
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)

	fmt.Print("Teks: ")
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}

	balikanArray(&tab, n)

	fmt.Print("Reverse teks: ")
	cetakArray(tab, n)
}