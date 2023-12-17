package main

import "fmt"

func main() {
	var p, d, k, tb int
	fmt.Scan(&p)
	if p <= 15 && p >= 1 {
		d = p / 5
		if d == 3 {
			d = 3
			fmt.Println("Dewasa :", d)
		} else {
			d = d + 1
			fmt.Println("Dewasa :", d)
		}
	} else if p >= 16 {
		d = 3
		p = p - 15
		if p <= 10 {
			k = p / 2
			fmt.Println("Dewasa :", d, "," , "Kecil :", k)
		} else {
			p = p - 10
			k = 5
			tb = p
			fmt.Println("Dewasa :", d, ",", "Kecil :", k , "," , "Tidak Berangkat :", tb)
		}
	} else {
		d = 0
		k = 0
		tb = 0
		fmt.Println("Tidak Ada")
	}
}
