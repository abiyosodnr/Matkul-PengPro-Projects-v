package main

import (
	"fmt"
	"sort"
)

func areDiceConsecutive(d1, d2, d3 int) bool {
	// Menggunakan slice untuk menyimpan mata dadu
	dice := []int{d1, d2, d3}

	// Mengurutkan mata dadu dalam urutan menaik
	sort.Ints(dice)

	// Memeriksa apakah selisih antara mata dadu adalah 1
	return dice[1]-dice[0] == 1 && dice[2]-dice[1] == 1
}

func main() {
	var d1, d2, d3 int

	fmt.Print("Masukkan mata dadu pertama: ")
	fmt.Scan(&d1)

	fmt.Print("Masukkan mata dadu kedua: ")
	fmt.Scan(&d2)

	fmt.Print("Masukkan mata dadu ketiga: ")
	fmt.Scan(&d3)

	if areDiceConsecutive(d1, d2, d3) {
		fmt.Println("Mata dadu muncul secara konsekutif.")
	} else {
		fmt.Println("Mata dadu tidak muncul secara konsekutif.")
	}
}
