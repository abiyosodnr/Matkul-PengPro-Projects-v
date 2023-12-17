package main

import "fmt"

func main() {

	var inp, hari, total int
	var rata2 float64
	var stop bool
	fmt.Scan(&inp)
	for !stop {
		// fmt.Scan(&inp)
		hari = hari + 1
		total = total + inp
		stop = inp == -1
	}
	rata2 = float64(total) / float64(hari)
	fmt.Println(fmt.Sprintf("%.2f", rata2))

}
