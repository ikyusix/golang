package main

import (
	"fmt"
)

func main() {
	var sisi, keliling, luas float32
	fmt.Print("Masukkan sisi sebuah persegi: ")
	fmt.Scan(&sisi)

	keliling = 4 * sisi
	luas = sisi * sisi

	fmt.Println("Hasil dari perhitungan luas dan keliling persegi dengan sisi:", sisi, "adalah sebagai berikut")
	fmt.Println("Persegi dengan sisi ", sisi, "mempunyai luas", luas, " dan mempunyai keliling", keliling)
}
