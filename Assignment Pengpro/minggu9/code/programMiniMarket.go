package main

import "fmt"

func main() {
	var n int
	var sedia, kartu, diskon, cashback bool
	fmt.Scan(&n, &sedia)
	kartu = sedia == true
	diskon = n >= 100000
	cashback = (n >= 200000) && (kartu == true)
	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)

}
