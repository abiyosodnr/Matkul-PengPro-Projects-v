package main

import "fmt"

func main() {
	// nomor 1
	// var i, n int
	// var in int
	// var hasil int
	// fmt.Scan(&n)
	// hasil = 0
	// for i = 1; i <= n; i = i + 1 {
	// 	fmt.Scan(&in)
	// 	hasil = hasil + in
	// }
	// fmt.Println(hasil)

	// nomor 2
	var i, n int
	var jam1, menitPuluhan1, menit1, jam2, menitPuluhan2, menit2 int
	var jumlah int
	var jumlahPerhari int
	var hasil bool
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&jam1, &menitPuluhan1, &menit1, &jam2, &menitPuluhan2, &menit2)
		jumlah = ((jam2 - jam1) * 60) + ((menitPuluhan2 - menitPuluhan1) * 10) + (menit2 - menit1)
		jumlahPerhari = jumlahPerhari + jumlah
	}
	hasil = jumlahPerhari >= (40 * 60)
	fmt.Println(hasil)
	fmt.Println(jumlah)
	fmt.Println(jumlahPerhari)
}
