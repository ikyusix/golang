package main

import "fmt"

func main() {

	a := 10
	b := 10 + 5 //aritnatik operation
	c := b - a
	d := b * c
	e := 7 % 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	//e++
	//e = e + 1
	e += 10 //assignment operation
	fmt.Println(e)

	condLeft := true // operation logic
	condRight := false

	condResult := condLeft || condRight

	fmt.Println("condLeft: ", condLeft)
	fmt.Println("condRight: ", condRight)
	fmt.Println("condResult: ", condResult)

	//operation realtional
	relasi := a > 20 || b < 80

	//compile go build
	//run go run
	//unsigned gada nol
	//realation bisa selain boolean untuk input ny
	//logical, inputnya sudah pasti boolean

	fmt.Println("relasi: ", relasi)
}
