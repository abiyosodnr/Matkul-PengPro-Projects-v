package main

import "fmt"

func main() {

	// soal 1

	// var a, b, c bool
	// fmt.Scan(&a, &b, &c)
	// if a == true && b == true && c == true {
	// 	fmt.Println("berkemah")
	// } else {
	// 	fmt.Println("tidak berkemah")
	// }

	// soal 2

	// soal 3

	// var a, b, c string
	// fmt.Scan(&a, &b, &c)
	// if a == b && b == c {
	// 	fmt.Println("imbang")
	// } else if a == c && b != c {
	// 	fmt.Println("pemain 2 pemenang")
	// } else if a == b && b != c {
	// 	fmt.Println("pemain 3 pemenang")
	// } else if b == c && a != b {
	// 	fmt.Println("pemain 1 pemenang")

	// }

	// soal 4

	// var a, b, c int
	// var x, y float64
	// fmt.Scan(&a)
	// fmt.Scan(&b)
	// fmt.Scan(&c)
	// a = a + 1
	// b = b + 1
	// x = float64(b) / float64(a)
	// y = float64(c) / float64(a)
	// if b+c <= a {
	// 	if x >= 0.6 {
	// 		fmt.Println("Kandidat A menang")
	// 	} else if y >= 0.6 {
	// 		fmt.Println("Kandidat B menang")
	// 	} else {
	// 		fmt.Println("Tidak ada pemenang")
	// 	}
	//}

	// QUIZ 3

	// var a, b, c, d, e, f, g int
	// fmt.Scan(&a, &b, &c, &d, &e, &f, &g)
	// if (a == 0 || a == 1) && (b == 0 || b == 1) && (c == 0 || c == 1) && (d == 0 || d == 1) && (e == 0 || e == 1) && (f == 0 || f == 1) && (g == 0 || g == 1) {
	// 	if g == 1 && a+b+c+d+e+f == 0 {
	// 		fmt.Println("ascending sorted")
	// 	} else if f+g == 2 && a+b+c+d+e == 0 {
	// 		fmt.Println("ascending sorted")
	// 	} else if e+f+g == 3 && a+b+c+d == 0 {
	// 		fmt.Println("ascending sorted")
	// 	} else if d+e+f+g == 4 && a+b+c == 0 {
	// 		fmt.Println("ascending sorted")
	// 	} else if c+d+e+f+g == 5 && a+b == 0 {
	// 		fmt.Println("ascending sorted")
	// 	} else if b+c+d+e+f+g == 6 && a == 0 {
	// 		fmt.Println("ascending sorted")
	// 	}
	// }

	// ASSESMENT

	// NO 1

	// var a, b int
	// fmt.Scan(&a, &b)
	// if a%2 == 0 && b%2 == 1 {
	// 	fmt.Println("berhasil")
	// } else if a%2 == 1 && b%2 == 0 {
	// 	fmt.Println("berhasil")
	// }

	// NO 2

	var j1, m1, lm, j2, m2 int
	fmt.Scan(&j1, &m1, &lm)
	j2 = j1
	m2 = m1 + lm
	if m2 >= 60 {
		j2 = j2 + (m2 / 60)
		m2 = m2 % 60
	}
	if j2 >= 20 {
		fmt.Println("balik lagi besok")
	} else {
		fmt.Println(j2, m2)
	}

	// No 3

	// var j1, j2 int
	// fmt.Scan(&j1)
	// if j1 >= 0 && j2 <= 23 {
	// 	if j1 == 0 {
	// 		j2 = 12
	// 		fmt.Println(j2, "AM")
	// 	} else if j1 >= 1 && j1 <= 11 {
	// 		j2 = j1
	// 		fmt.Println(j2, "AM")
	// 	} else if j1 == 12 {
	// 		j2 = j1
	// 		fmt.Println(j2, "PM")
	// 	} else if j1 >= 13 && j1 <= 23 {
	// 		j2 = j1 - 12
	// 		fmt.Println(j2, "PM")
	// 	}
	// }

	// var x,y,z int

	// fmt.Scan(&x,&y,&z)

	// if (x/100) * z <= z {
	// 	fmt.Println(z-y * (x/100))
	// }else if (x/100) * z > y {
	// 	fmt.Println(z - y)
	// }

}
