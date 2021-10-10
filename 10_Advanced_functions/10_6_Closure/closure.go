package main

import "fmt"


func closure() func(){
	text := "inside Closure"

	func1 := func ()  {
		fmt.Println(text)
	}

	return func1
}

func main()  {
	text := "inside main"
	fmt.Println(text)

	newFunc := closure()
	newFunc()
	
}