package main

import "fmt"

func main() {
	var i, hari, jam, jumlah int
	var rata2 float64
	fmt.Scan(&hari)
	for i = 1; i <= hari; i++ {
		fmt.Scan(&jam)
		jumlah = jumlah + jam
	}
	rata2 = float64(jumlah) / float64(hari)
	fmt.Println(rata2)
}
