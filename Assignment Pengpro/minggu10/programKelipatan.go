package main

import "fmt"

func main() {
	var bil int
	fmt.Scan(&bil)
	if bil%3 == 0 || bil%5 == 0 {
		if bil%3 == 0 {
			fmt.Println("Kelipatan 3")
		}
		if bil%5 == 0 {
			fmt.Println("Kelipatan 5")
		}
	}
}
