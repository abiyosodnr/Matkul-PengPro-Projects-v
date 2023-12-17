package main

import "fmt"

func main() {
	var jb string
	var mk, ja, gaji, tja, hasil int
	fmt.Scan(&jb, &mk, &ja)
	if ja >= 3 {
		ja = 3
	}
	if jb == "Staf" {
		if mk < 5 {
			gaji = 4000
		} else if mk > 10 {
			gaji = 5000
		} else if mk >= 5 && mk <= 10 {
			gaji = 4000
		}
		if mk >= 5 {
			tja = ja * 100
		}
	} else if jb == "Manager" {
		if mk > 10 {
			gaji = 10000
		} else if mk <= 10 {
			gaji = 8500
		}
		tja = ja * 300
	} else if jb == "Direktur" {
		gaji = 20000
		tja = ja * 500
	}
	hasil = gaji + tja
	fmt.Println(gaji, "+", tja, "=", hasil)
}
