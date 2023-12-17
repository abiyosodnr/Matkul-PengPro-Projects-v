package main

import "fmt"

func main() {
	// var x, y int
	// fmt.Scan(&x, &y)
	// for x <= y {
	// 	fmt.Println(x)
	// 	x += 1

	// }
	

	var n int
	var i int
	var sisi, k, l int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&sisi)
		l = sisi * sisi
		k = sisi * 4
		fmt.Println(l, k)
	}

}
