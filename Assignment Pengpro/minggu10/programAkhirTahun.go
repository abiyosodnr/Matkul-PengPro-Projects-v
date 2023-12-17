package main

import "fmt"

func main() {
	var x int
	var kartu, diskon, cashback bool
	fmt.Scan(&x, &kartu)
	if x >= 100000 {
		diskon = true
		x = x - (x * 10 / 100)
	}
	if x >= 200000 && kartu == true {
		cashback = true
		x = x - 75000
	}
	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
	fmt.Println("Rp.", x)

}
