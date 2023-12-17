package main

import "fmt"

func main() {
	var th int
	var hasil bool
	fmt.Scan(&th)
	if th % 400 == 0 || (th % 4 == 0 && th % 100 != 0) {
		hasil = true
	}else {
		hasil = false
	}
	fmt.Println(hasil)
}