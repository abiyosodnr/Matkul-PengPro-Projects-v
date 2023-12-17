package main

import "fmt"

func main() {
	var i, n, total int
	i = 1
	n = 91
	total = 0
	for i = 1; i <= n; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			total = total + 10000
		}else if i % 3 == 0 {
			total = total + 5000
		}else if i % 5 == 0 {
			total = total + 8000
		}else {
			total = total + 1000
		} 
	}
	fmt.Println(total)
}