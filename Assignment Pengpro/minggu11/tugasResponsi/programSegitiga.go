package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a + b > c && a + c > b && b + c > a {
		if (a == b && a != c) || (a == c && a != b) || (b == c && a != b) {
			fmt.Println("segitiga sama kaki")
		}else if a == b && b == c && a == c {
			fmt.Println("segitiga sama sisi")
		}else if a*a + b*b == c*c || b*b + c*c == a*a || a*a + c*c == b*b {
			fmt.Println("segitiga siku siku")
		}else {
			fmt.Println("segitiga bebas")
		}
	}
}