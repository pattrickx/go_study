package main

import "fmt"

func main()  {
	var text string = "variable1"
	text2 := "Variable2"
	fmt.Println(text)
	fmt.Println(text2)
	var (
		text3 string = "text3"
		text4 string = "text4"
	)
	fmt.Println(text3,text4)

	text5, text6 := "text5", "text6"
	fmt.Println(text5,text6)

	text5, text6 = text6, text5
	fmt.Println(text5,text6)
	
	const const1 string = "const1"
	fmt.Println(const1)
}