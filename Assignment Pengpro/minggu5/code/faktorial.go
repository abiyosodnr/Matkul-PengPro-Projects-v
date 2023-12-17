package main

import "fmt"

func main() {
	var i, n int
	var hasil int
	hasil = 1
	fmt.Scan(&n)
	for i = n; i >= 1; i-- {
		hasil = hasil * i
	}
	fmt.Println(hasil)
}
