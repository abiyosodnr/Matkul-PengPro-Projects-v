package main

import "fmt"

func main() {
	var n byte
	var hasil bool
	fmt.Scanf("%c", &n)
	hasil = (n >= 'a' && n <= 'z') || (n >= '0' && n <= '9') || (n >= 'A' && n <= 'Z')
	fmt.Println(hasil)
}
