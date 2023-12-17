package main

import "fmt"

func main() { //start format

	nama := "Abiyoso Danar Panji Yudhanto" //string type
	// benar := true                          //boolean type
	// age := 18                              //integer type
	var nomorHuruf = nama[0]
	var nomorHurufstring = string(nomorHuruf) // konversi data byte ke string

	fmt.Println("Benar =", true)
	fmt.Println("Salah =", false)
	fmt.Println(nama)
	fmt.Println(nomorHurufstring)

	fmt.Println("Abiyoso Danar")
	fmt.Println(len("Abiyoso Danar"))
	fmt.Println(len("aku mau ayam"))

	var ( //multiple variable
		rumah   = "kursi, meja, kamar"
		sekolah = "meja, guru"
		// ruangKelas   = 13
		// umurRataRata = 46.5
	)
	fmt.Println(sekolah)
	fmt.Println(rumah)

	var age int16 = 127 // conversi data
	var age2 = int8(age)

	fmt.Println(age2)

	const ( //constan
		SMA    = "5 Depok"
		Negara = "Indonesia"
	)
	fmt.Println(SMA)

	var a = 100 //operasi matematika
	var b = 20
	var (
		c = a + b
		d = a - b
		e = a * b
		f = a / b
		g = a % b
	)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	var i = 20
	i += 30 //shortcut operasi matematika
	i++     // shortcut unary operator
	fmt.Println(i)

	var (
		value1 = 20
		value2 = 80
	)
	fmt.Println(value1 == value2)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)
	fmt.Println(value1 != value2) //operasi perbandingan

	var nilaiFisika = 81
	var nilaiMatematika = 73

	var lulusFisika = nilaiFisika >= 75
	var lulusMatematika = nilaiMatematika >= 75

	var lulus = lulusFisika || lulusMatematika //operasi boolean
	fmt.Println(lulus)

	fmt.Println(nilaiFisika >= 75 && nilaiMatematika >= 75) //shortcut operasi boolean

	var namaNama = [...]string{ //array
		"abi",
		"anwar",
		"fadly",
		"ariza",
		"falah",
	}
	fmt.Println(namaNama)

}
