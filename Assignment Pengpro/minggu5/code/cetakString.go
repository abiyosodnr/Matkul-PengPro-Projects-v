package main

import "fmt"

func main() {
	var i, n int
	var word string
	fmt.Scan(&n)
	fmt.Scan(&word)
	for i = 1; i <= n; i++ {
		fmt.Println(word)
	}
}
