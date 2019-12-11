package main

import "fmt"

func main() {

	a, b := 10, 20

	a, b = b, a

	fmt.Println(a, b)

	// var a = 10
	// var b = 20

	// c := a
	// a = b
	// b = c

	// fmt.Println(a, b)

}
