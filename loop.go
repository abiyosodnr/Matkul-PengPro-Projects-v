package main

import "fmt"

func main() {
	var in1, in2, in3 int
	var x string
	fmt.Scan(&in1, &in2, &in3)
	if (in1+in2+in3)%2 == 0 {
		x = "genap"
	} else {
		x = "ganjil"
	}

	fmt.Println(x)
}
