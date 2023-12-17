package main

import "fmt"

func main() {
	var x int
	var d1, d2, d3 int

	fmt.Scan(&x)
	d1 = x / 100
	x = x % 100
	d2 = x / 10
	x = x % 10
	d3 = x / 1

	fmt.Println(d1, d2, d3)
}

// program TigaDigit
// kamus
//     x : integer
//     d1, d2, d3 : integer
// algoritma
//     input(x)
//     d1 <- x div 100
//     x <- x mod 100
//     d2 <- x div 10
//     x <- mod 10
//     d3 <- x mod 1
//     output(d3, d2, d1)
// endprogram
