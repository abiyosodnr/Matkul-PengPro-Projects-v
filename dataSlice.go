package main

import "fmt"

func main() {

	var urutan = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var sliceurutan1 = urutan[:3]
	urutan[1] = 890
	var tambahurutan = append(sliceurutan1, 3, 4, 5)
	var a = len(sliceurutan1)
	var b = cap(sliceurutan1)

	fmt.Println(sliceurutan1)
	fmt.Println("yauda lah ya")
	fmt.Println(tambahurutan)
	fmt.Println(a, b)
}
