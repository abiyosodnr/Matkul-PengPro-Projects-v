package main

import "fmt"

func main() {
	var h1, m1, h2, m2, h, m int
	fmt.Scan(&h1, &m1, &h2, &m2)
	if h1 > h2 || m1 > m2 {
		if h1 > h2 {
			h2 = h2 + 12
		}
		if m1 > m2 {
			m2 = m2 + 60
			h2 = h2 - 1
		}
	}
	h = h2 - h1
	m = m2 - m1
	fmt.Println(h, "jam", m, "menit")

}
