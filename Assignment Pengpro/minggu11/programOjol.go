package main

import "fmt"

func main() {
	var jam, mnt, tarif int
	var jrk float64
	fmt.Scan(&jam, &mnt, jrk)
	if (jam >= 5 && jam <= 22 && mnt == 0) && jrk <= 20 {
		if (jam >= 5 && (jam <= 9 && mnt == 0)) || (jam >= 16 && (jam <= 19 && mnt == 0)) {
			if jrk > 0 && jrk <= 10 {
				tarif = int(jrk) * 5000
			} else if jrk > 10 && jrk <= 20 {
				tarif = int(jrk) * 4500
			}
		} else {
			tarif = int(jrk) * 4000
		}
		fmt.Println(tarif)
	} else {
		fmt.Println("Maaf, kami tidak bisa melayani pesanan anda.")
	}
}
