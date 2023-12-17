package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a != b && a != c && b != c {
		fmt.Println("Segitiga Sembarang")
	} else if a == b && a == c && b == c {
		fmt.Println("Segitiga Sama Sisi")
	} else if (a == c && a != b) || (a == b && a != c) || (b == c && a != b) {
		fmt.Println("Segitiga Sama Kaki")
	}
}
