package main

import "fmt"

// variable global disimpannya di luar func
// dan usahakan hanya cont yg global, sisanya variable local

//main function yang pertamakali dijalankan
func main() {
	// \n untuk enter

	// pendeklarasian variable bisa lebih dari satu variabel.
	// var angka, angka2 int = 50, 60
	// angka, angka2 := 700, 800
	/*
		var (
			angka int = 9
			angka2 int = 10
		)
	*/

	var angka int8 // -125 s/d 126 angka yang bisa ditampung
	// = itu operator asignment
	angka = 100

	var kal string

	kal = "ini adalah sebuah string"

	var koma float32
	// untuk menampung angka berkoma
	// googling float 32 bisa menampung berapa, 64 berapa
	koma = 1.5

	var bol bool
	// untuk mengisi logic d komputer. benar atau salah
	bol = true
	bol = false
	// diprogramming itu berjalan seraca sequencial atau berurutan

	// cara penulisan variable
	angka3 := 10
	angka2 := angka // asignment dengan variable lain

	kalimat2 := "kalimat 2"

	float2 := 6.5

	bol2 := true

	// cara penulisan variable yang ketiga
	const phi = 3.14
	// const adalah nilai tetap yang tidak boleh dirubah. phi, gravitasi, conversi celcius ke kelvin

	var x2 string = "bima"
	var y string = "adam"
	fmt.Println(x2 == y)

	var x string
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)

	fmt.Println(angka)
	fmt.Println(angka2)
	fmt.Println(angka3)
	fmt.Println(kal)
	fmt.Println(kalimat2)
	fmt.Println(koma)
	fmt.Println(float2)
	fmt.Println(bol)
	fmt.Println(bol2)
	fmt.Println(phi)

}
