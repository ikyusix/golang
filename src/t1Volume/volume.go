package main

import (
	"fmt"
)

func main() {
	var panjang, lebar, tinggi, volume, luasPermukaan int
	fmt.Print("Masukkan panjang: ")
	fmt.Scan(&panjang)

	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&lebar)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&tinggi)

	volume = panjang * lebar * tinggi
	luasPermukaan = 2 * ((panjang * lebar) + (panjang * tinggi) + (lebar * tinggi))

	fmt.Println("Hasil dari perhitungan luas permukaan dan volume dari sebuah balok adalah sebagai berikut:")
	fmt.Println("Balok dengan panjang:", panjang)
	fmt.Println("Lebar:", lebar)
	fmt.Println("Dan tinggi:", tinggi)
	fmt.Println("memiliki luas permukaan:", luasPermukaan, "dan volume:", volume)
}
