package main

import (
	"fmt"
	"strings"
)

// func main() {

	// var i int = 6
	// var f float32 = 10.5

	// var s string
	// fmt.Println(100, "ini test")
	// s = fmt.Sprintln(100, "ini test") // semuanya dirubah kedalah string

	// fmt.Println("")

	// fmt.Println("isi variable i adalah ", s)

	// fmt.Println("abis")
	//======================================================================================
	//print f
	// a, b := 100, 200
	// fmt.Printf("printf: ini %%v dan ini %v", a, b)

	//======================================================================================
	//flow control, condition, loop, defer
	// var a int = 5

	// if a < 10 && a >= 8 {
	// 	fmt.Println("perintah dieksekusi")
	// } else if a == 7 {
	// 	fmt.Println("else if yang dijalankan")
	// } else if a > 9 {
	// 	fmt.Println("else if ke dua")
	// } else {
	// 	fmt.Println("else terakhir")
	// }

	//for
	// for i := 1; i < 10; i++ {
	// 	fmt.Println("bima", i)
	// }

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	//defer
	// defer fmt.Println("halo")
	// fmt.Println("Selamat datang")

	//======================================================================================
	//Array
	// var names [3]string
	// names[0] = "Bima"
	// names[1] = "Adam"
	// names[2] = "Nugraha"

	// fmt.Println(names[0], names[1], names[2])

	// var names [6]string
	// names[0] = "Bima"
	// names[1] = "Adam"
	// names[2] = "Nugraha"
	// names[3] = "Nugraha"
	// names[4] = "Nugraha"
	// names[5] = "Nugraha"

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// var names [3]string
	// names[0] = "Bima"
	// names[1] = "Adam"
	// names[2] = "Nugraha"

	// var fruits = [3]string{"nanas", "semangka", "buah weh"}

	// var num = [...]int{1,2,3,4,5,6,7,8}

	// fmt.Println(names[0], names[1], names[2])
	// fmt.Println(fruits)
	// fmt.Println(num)

	// name := " Bang dani bawa makan"
	// fmt.Println(name)
	// fmt.Println(strings.TrimSpace(name))

	// 	var names = []string{"John", "Wick"}
	// 	printMessage("halo", names)
	// }

	// func printMessage(message string, arr []string) {
	// 	var nameString = strings.Join(arr, " ")
	// 	fmt.Println(message, nameString)

// }


//multiple function
func main() {
	name := "        Bang dani bawa makan"
	alamat := "   alamat"
	cutTheStrings(name, alamat)
}

func cutTheStrings(lu, gue string) {
	fmt.Println(strings.TrimSpace(lu))
	fmt.Println(strings.TrimSpace(gue))
}