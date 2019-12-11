package main

import (
	"fmt"
)

const phi = 3.14

func main() {
	var jariJari, keliling, luas float32
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&jariJari)

	keliling = 2 * phi * jariJari
	luas = phi * jariJari * jariJari

	fmt.Println("Hasil dari perhitungan luas dan keliling lingkaran dengan jari-jari:", jariJari, "adalah sebagai berikut")
	fmt.Println("Lingkaran dengan jarijari ", jariJari, "mempunyai luas", luas, " dan mempunyai keliling", keliling)
}
