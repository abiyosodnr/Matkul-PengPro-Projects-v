package main

import "fmt"

func main() {
	// var x, y int
	// fmt.Scan(&x, &y)
	// for x <= y {
	// 	fmt.Println(x)
	// 	x += 1

	// }

	// var n int
	// var i int
	// var sisi, k, l int
	// fmt.Scan(&n)
	// for i = 1; i <= n; i++ {
	// 	fmt.Scan(&sisi)
	// 	l = sisi * sisi
	// 	k = sisi * 4
	// 	fmt.Println(l, k)
	//}

	// var n int
	// var i int
	// var jumlah, jam, latihan int
	// var rata2 float64
	// fmt.Scan(&n)
	// jumlah = 0
	// for i = 0; i < n; i++ {
	// 	fmt.Scan(jam)
	// 	latihan = jumlah + jam
	// }
	// rata2 = float64(latihan) / float64(n)
	// fmt.Println(rata2)

	// var i, banyak int
	// var latihan float64
	// var jumlah float64

	// fmt.Scan(&banyak)
	// jumlah = 0
	// for i = 1; i <= banyak; i++ {
	// 	fmt.Scan(&latihan)
	// 	jumlah = jumlah + latihan
	// }
	// fmt.Println(float64(jumlah) / float64(banyak))

	// var in int
	// var i int
	// fmt.Scan(&in)
	// for i = 1; i <= in; i++ {
	// 	fmt.Println(i, i%2 == 0)
	// }

	// var i, hari, latihan, jumlah int
	// var hasil float64

	// fmt.Scan(&hari)
	// for i = 0; i < hari; i++ {
	// 	fmt.Scan(&latihan)
	// 	jumlah += latihan

	// }
	// hasil = float64(jumlah) / float64(hari)
	// fmt.Println(hasil, latihan)

	// var i, n int
	// var hasil int
	// fmt.Scan(&n)
	// hasil = 1
	// for i = n; i > 0; i-- {
	// 	hasil *= i
	// }
	// fmt.Println(hasil)

	var i, n int
	var hasil bool
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		hasil = n%i == 0
		fmt.Println(i, hasil)
	}

}
