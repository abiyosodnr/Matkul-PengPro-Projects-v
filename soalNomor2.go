package main

import "fmt"

func main() {
	var x float64 // nilai input integer
	var f float64 // output float

	fmt.Println("Masukkan Nilai x")
	fmt.Scan(&x)

	f = (x*x*x + 3*x) / (x*x*x*x - 3*x*x + 4)

	fmt.Println(f)

}

// program MatematikaB
// kamus
//     x : integer
//     f : real
// algoritma
//     input(x)
//     f <- (x*x*x + 3*x)/(x*x*x*x-3*x*x+4)
//     output(f)
// endprogram
