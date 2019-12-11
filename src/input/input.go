package main

import (
	"fmt"
)

func main() {
	var angka1, angka2, hasil int
	fmt.Print("Masukkan angka 1: ")
	fmt.Scan(&angka1)

	fmt.Print("Masukkan angka 2: ")
	fmt.Scan(&angka2)

	hasil = angka1 + angka2
	fmt.Println("hasil penjumlahan adalah ", hasil)

	// var kal string
	// fmt.Print("Masukkan kalimat: ")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// kal = scanner.Text()
	// fmt.Println("Kalimat yang dimasukkan adalah: ", kal)

}
