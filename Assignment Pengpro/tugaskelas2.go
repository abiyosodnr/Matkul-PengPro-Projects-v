package main

import "fmt"

func main() {
	// var i, n int
	// var jumlah int

	// i = 2
	// n = 50
	// for i = 2; i <= n; i++ {
	// 	fmt.Println(i)
	// 	jumlah = jumlah + i
	// }
	// fmt.Println(i)
	// fmt.Println(jumlah)

	// var i, n int
	// var kata string
	// fmt.Scan(&n)
	// fmt.Scan(&kata)
	// for i = 1; i <= n; i++ {
	// 	fmt.Println(kata)
	// }

	var i, n int
	var teks, gambar int
	var jumlteks, jumlgambar int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&teks, &gambar)
		jumlteks = jumlteks + teks
		jumlgambar = jumlgambar + gambar
	}
	fmt.Println(jumlteks, jumlgambar)
}
