package main

import "fmt"

func main() {
	// var a, b, c, hasil int
	// fmt.Scan(&a)
	// fmt.Scan(&b)
	// fmt.Scan(&c)
	// fmt.Scan(&hasil)
	// if a+b+c == hasil {
	// 	fmt.Println("benar")
	// } else if a+b+c != hasil {
	// 	fmt.Println("salah")
	// }

	// var rank int
	// fmt.Scan(&rank)
	// if rank <= 5 && rank >= 1 {
	// 	fmt.Println("mendapat hadiah")
	// }

	// var a, b, c, d int
	// var rata2 float64
	// fmt.Scan(&a, &b, &c, &d)
	// rata2 = (float64(a) + float64(b) + float64(c) + float64(d)) / 4
	// if rata2 >= 3.5 {
	// 	fmt.Println("bagus")
	// } else if rata2 <= 1.5 {
	// 	fmt.Println("kurang")
	// } else {
	// 	fmt.Println("sedang")
	// }

	// var a, b, c string
	// fmt.Scan(&a, &b, &c)
	// if a == b && a != c {
	// 	fmt.Println("pemain 3 pemenang")
	// } else if a == c && a != b {
	// 	fmt.Println("pemain 2 pemenang")
	// } else if b == c && a != b {
	// 	fmt.Println("pemain 1 pemenang")
	// } else if a == b && a == c && b == c {
	// 	fmt.Println("imbang")
	// }

	var a, b bool
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a == true && b == true {
		fmt.Println("keluar jalan-jalan")
	} else if (a == false && b == true) || (a == true && b == false) {
		fmt.Println("diam di rumah saja")
	} else if a == false && b == false {
		fmt.Println("diam di rumah saja")

	}

	// var a, b int
	// var per float64

	// fmt.Scan(&a)
	// fmt.Scan(&b)
	// per = float64(b) / float64(a)
	// fmt.Println(per)
	// if per >= 0.5 && per <= 0.75 {
	// 	fmt.Println("berangkat")
	// }else {
	// 	fmt.Println("tidak berangkat")
	// }

	// var a, b, c, d, ganjil, genap int

	// fmt.Scan(&a, &b, &c, &d)
	// if a % 2 == 0 {
	// 	genap = genap + 1
	// }else {
	// 	ganjil = ganjil + 1
	// }
	// if b % 2 == 0 {
	// 	genap = genap + 1
	// }else {
	// 	ganjil = ganjil + 1
	// }
	// if c % 2 == 0 {
	// 	genap = genap + 1
	// }else {
	// 	ganjil = ganjil + 1
	// }
	// if d % 2 == 0 {
	// 	genap = genap + 1
	// }else {
	// 	ganjil = ganjil + 1
	// }
	// fmt.Println(genap, ganjil)

	// var a, b, c int

	// fmt.Scan(&a, &b, &c)
	// if a < b && a < c {
	// 	fmt.Println(a)
	// }else if b < a && b < c {
	// 	fmt.Println(b)
	// }else if c < a && c < b {
	// 	fmt.Println(c)
	// }

	// var a, b, c int

	// fmt.Scan(&a, &b, &c)
	// if a < b && b < c {
	// 	fmt.Println(c, b, a)
	// } else if b < a && a < c {
	// 	fmt.Println(c, a, b)
	// } else if c < a && a < b {
	// 	fmt.Println(b, a, c)
	// } else if a < c && c < b {
	// 	fmt.Println(b, c, a)
	// } else if b < c && c < a {
	// 	fmt.Println(a, c, b)
	// } else if c < b && b < a {
	// 	fmt.Println(a, b, c)
	// } else if a == b && c < a {
	// 	fmt.Println(a, b, c)
	// } else if b == c && a < b {
	// 	fmt.Println(b, c, a)
	// } else if a == c && b < a {
	// 	fmt.Println(a, c, b)
	// } else if a == b && a == c && b == c {
	// 	fmt.Println(a, b, c)
	// } else if a == b && c > a {
	// 	fmt.Println(c, a, b)
	// } else if a == c && b > a {
	// 	fmt.Println(b, a, c)
	// } else if b == c && a > b {
	// 	fmt.Println(a, b, c)
	// }

}
