package main

import "fmt"

func main() {
	var i, n int
	var s bool
	fmt.Scan(&n)
	for i = 1; i <= n; i += 1 {
		s = n%i == 0
		fmt.Println(i, s)
	}
}
