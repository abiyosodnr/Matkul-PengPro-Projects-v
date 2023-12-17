package main

import "fmt"

func main() {
	var nama, asal, prodi string
	var NIM uint64

	nama = "Abiyoso Danar Panji Yudhanto"
	prodi = "S1 Informatika"
	NIM = 103012300006
	asal = "Depok"

	const (
		Depok   = "Sawangan"
		KodePos = 16358
	)

	var (
		age  int16 = 127		// conversi data
		age2       = int8(age)
	)

	fmt.Println("halo nama saya ", nama, "saya dari prodi", prodi, "dengan NIM", NIM, "berasal dari kota", asal)
	fmt.Println(KodePos)
	fmt.Println(age, age2)
}
