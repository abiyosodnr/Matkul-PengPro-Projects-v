package main

import "fmt"

func main() {
	var mh, skr  bool
	var lm int
	var total float64
	
	fmt.Scan(&mh, &skr, &lm, &total)
	if mh == true {
		if lm > 12 {
			fmt.Println("Pelanggan Prioritas I")
			if skr == true {
				total = total * 0.65
			}else {
				total = total * 0.7
			}
		}else if lm >= 8 && lm <= 12 {
			fmt.Println("Pelanggan Prioritas II")
			if skr == true {
				total = total * 0.8
			}else {
				total = total * 0.85
			}
		}else {
			total = total * 0.9
		}
		
	}
	fmt.Println(total)
}