package main

import "fmt"

func main() {
	var x byte
	var hasil bool
	fmt.Scanf("%c", &x)
	if (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z') || (x >= '1' && x <= '9' ) {
		hasil = true
	}else {
		hasil = false
	}
	fmt.Println(hasil)
}