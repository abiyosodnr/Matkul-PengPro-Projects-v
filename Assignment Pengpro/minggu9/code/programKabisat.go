package main

import "fmt"

func main() {
	var n int
	var hasil bool
	fmt.Scan(&n)
	hasil = (n%400 == 0) || ((n%4 == 0) && (n%100 != 0))
	fmt.Println(hasil)
}
