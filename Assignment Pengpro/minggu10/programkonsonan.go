package main

import "fmt"

func main() {
	var huruf byte
	fmt.Scanf("%c", &huruf)

	if huruf > 'a' && huruf <= 'z' && huruf != 'i' && huruf != 'u' && huruf != 'e' && huruf != 'o' {
		fmt.Println("konsonan")
	} else if huruf > 'A' && huruf <= 'Z' && huruf != 'I' && huruf != 'U' && huruf != 'E' && huruf != 'O' {
		fmt.Println("konsonan")

	} else {
		fmt.Println("bukan konsonan")
	}
}
