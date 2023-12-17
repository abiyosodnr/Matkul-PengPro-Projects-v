package main

import "fmt"

func main() {
	var a, b, c int
	var hasil bool
	fmt.Scan(&a, &b, &c)
	hasil = ((a+b)/2 == c) || ((b+c)/2 == a) || ((a+c)/2 == b)
	fmt.Println(hasil)
}
