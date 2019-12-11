package main

import (
	"fmt"
)

func main() {
	var i int = 10
	var f float32 = 10.5

	fmt.Println(float32(i) + f)

	var z int = 65

	fmt.Println(string(z))

	//Sprintln -> mengkonvert int ke string

	var s string
	s = fmt.Sprintln(i, "ini juga diprint")
	fmt.Println("isi variabel i adalah ", s)

	fmt.Printf("ini %[2]v adalah fungsi %[1]v printf %%V", s, 99)
}
