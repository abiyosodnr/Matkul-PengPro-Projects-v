package main

import "fmt"

func main() {
	var k1, k2, amount float64
	fmt.Scan(&k1, &k2)
	if k1 < k2 {
		amount = k2 - k1
		fmt.Println("Peningkatan sebesar", amount)
	} else if k1 > k2 {
		amount = k1 - k2
		fmt.Println("Penurunan sebesar", amount)
	} else if k1 == k2 {
		fmt.Println("Tetap")
	}
}
