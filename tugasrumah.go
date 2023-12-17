package main

import "fmt"

func main() {
	// const pi float64 = 3.14
	// var r, t int
	// var hasil float64

	// fmt.Scan(&r, &t)

	// hasil = (1.0 / 3.0) * pi * float64(r) * float64(r) * float64(t)

	// fmt.Println(hasil)

	// var adik, kakak, x int
	// var hasil bool

	// fmt.Scan(&adik, &kakak)

	// x = adik - kakak

	// hasil = x%2 == 1

	// fmt.Println(hasil)

	// var b1, b2, b3, x int
	// var kg, g int

	// fmt.Scan(&b1, &b2, &b3)

	// x = b1 + b2 + b3

	// kg = x / 1000
	// g = x % 1000

	// fmt.Println(kg, g)

	// var p, q, r, s int
	// var t int

	// fmt.Scan(&p)

	// q = (p % 100 % 10 * 100)
	// r = (p % 100 / 10 * 10)
	// s = (p / 100)

	// t = q + r + s

	// fmt.Println(t)

	// const pi float64 = 3.14
	// var r float64
	// var v float64

	// fmt.Scan(&r)

	// v = 4.0 / 3.0 * pi * r * r * r

	// fmt.Println(v)

	// var adik, kakak int
	// var selisih, selisih2, selisihfix, hasil bool

	// fmt.Scan(&adik, &kakak)
	// selisih = adik-kakak == 3 || adik-kakak == 6
	// selisih2 = adik-kakak == -3 || adik-kakak == -6
	// selisihfix = selisih || selisih2
	// hasil = adik == kakak || selisihfix

	// fmt.Print(hasil)

	var in1, in2 string
	var a, x, y, z byte
	var hasil bool

	fmt.Scan(&in1, &in2)

	a = in1[0] - 32
	x = in1[0] + 32
	y = in2[0]
	z = in1[0]

	hasil = x == y || z == y || a == y && z >= 65 || z <= 90 && 

	fmt.Println(hasil)
}
