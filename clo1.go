package main

import "fmt"

func main() {
	// var input int
	// var jam, menit, detik int

	// fmt.Scan(&input)

	// jam = input / 3600
	// input = input % 3600
	// menit = input / 60
	// input = input % 60
	// detik = input / 1

	// fmt.Println(jam, menit, detik)

	// var a, b int
	// var x float64

	// fmt.Scan(&a, &b)

	// x = float64(a) / float64(b)

	// fmt.Println(x)

	// var w, x, y, z float64
	// var output float64

	// fmt.Scan(&w, &x, &y, &z)

	// output = (1.0 + 3.0*w*w) / (4.0*x - 5.0*y*y*y + 6.0*z*z*z*z)

	// fmt.Println(output)

	// var x float64
	// var y float64

	// fmt.Scan(&x)

	// y = ((x * x * x) + (3.0 * x)) / ((x * x * x * x) - (3 * x * x) + 4.0)

	// fmt.Println(y)

	// var i, n int
	// fmt.Scan(&i, &n)

	// for i = i; i <= n; i++ {

	// 	fmt.Printf("%d", i)
	// }

	// var n, m, i int
	// fmt.Scan(&n, &m)
	// for i = n; i <= m; i++ {
	// 	fmt.Printf("%d", i)
	// }

	// var i, n int
	// var in, jmldigit, jumlah int
	// fmt.Scan(&n)
	// for i = 1; i <= n; i++ {
	// 	fmt.Scan(&in)
	// 	jmldigit = (in / 1000) + (in % 10)
	// 	jumlah = jumlah + jmldigit
	// }
	// fmt.Println(jumlah)

	// var i, n int
	// var in string
	// fmt.Scan(&n)
	// fmt.Scan(&in)
	// for i = 1; i <= n; i++ {
	// 	fmt.Println(in)
	// }

	// var i, n int
	// var sisi int
	// var luas, keliling int
	// fmt.Scan(&n)
	// for i = 1; i <= n; i++ {
	// 	fmt.Scan(&sisi)
	// 	luas = sisi * sisi
	// 	keliling = sisi * 4
	// 	fmt.Println(luas, keliling)
	// }

	// var i, n int
	// var jmlhari, jumlah int
	// var rata2 float64
	// fmt.Scan(&n)
	// for i = 1; i <= n; i = i + 1 {
	// 	fmt.Scan(&jmlhari)
	// 	jumlah = jumlah + jmlhari
	// }
	// rata2 = float64(jumlah) / float64(n)
	// fmt.Println(rata2)

	// var i, n int
	// var hasil int
	// fmt.Scan(&n)
	// hasil = 1
	// for i = n; i >= 1; i-- {
	// 	fmt.Println(i)
	// 	hasil = hasil * i

	// }
	// fmt.Println(hasil)

	// var i, n int
	// var hasil bool
	// fmt.Scan(&n)
	// for i = 1; i <= n; i++ {
	// 	hasil = n%i == 0
	// 	fmt.Println(i, hasil)
	// }

	// var masukan string
	// var x byte
	// fmt.Scan(&masukan)
	// x = masukan[0]
	// fmt.Println(x)

	// var i, n int
	// var jam, menit, hasilJam, JamdrMenit, hasilMenit int
	// fmt.Scan(&n)
	// for i = 1; i <= n; i++ {
	// 	fmt.Scan(&jam, &menit)
	// 	JamdrMenit = JamdrMenit + menit
	// 	hasilJam = hasilJam + jam

	// }
	// hasilJam = hasilJam + (JamdrMenit / 60)
	// hasilMenit = JamdrMenit % 60
	// fmt.Println(hasilJam, hasilMenit)

	// var i, n int
	// var menit int
	// var hasiljam, hasilmenit int
	// fmt.Scan(&n)
	// for i = 1; i <= n; i++ {
	// 	fmt.Scan(&menit)
	// 	hasilmenit = hasilmenit + menit
	// }
	// hasiljam = hasilmenit / 60
	// hasilmenit = hasilmenit % 60
	// fmt.Println(hasiljam, hasilmenit)

	var i, n int
	var tabungan1, tabungan2, tabungan3 int
	fmt.Scan(&tabungan1, &n)
	tabungan2 = tabungan1
	tabungan3 = tabungan1
	for i = 2; i <= n; i++ {
		tabungan2 = tabungan2 + 2500
		tabungan3 = tabungan3 + tabungan2
	}
	tabungan2 = tabungan2 + (2500 * n)
	fmt.Println(tabungan2)
	fmt.Println(tabungan3)

	// var i, n int
	// var pendapatan, bagipendapatan, bagiaset, totalpendapatan, totalaset int
	// fmt.Scan(&n)
	// for i = 1; i <= n; i++ {
	// 	fmt.Scan(&pendapatan)
	// 	bagipendapatan = pendapatan / 3
	// 	bagiaset = pendapatan % 3
	// 	totalpendapatan = totalpendapatan + bagipendapatan
	// 	totalaset = totalaset + bagiaset
	// }
	// fmt.Println(totalpendapatan, totalaset)

}
