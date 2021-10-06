package main

import {
	"fmt"
}

func main() {
	var num int16 = 10000
	fmt.Println(num)

	var num1 int32 = 10000000
	fmt.Println(num1)

	//alias
	//INT32 = RUNE
	var num3 rune = 123456
	fmt.Println(num3)

	//BYTE = UINT8
	var num4 int16 = 123
	fmt.Println(num4)

	// FLOAT
	var numf1 float32 = 123000000.58
	var numf2 float64 = 123000.58

	fmt.Println(numf1, numf2)

	//char
	char1 := 'A' 
	fmt.Println(char1)

	var text int8
	fmt.Println(text)

	// ERROS
	var erro1 error = errors.New("Erroooo")
	fmt.Println(erro1)

}
