package main

import "fmt"


func main()  {
	fmt.Println("anonymous function")

	r := func(t string)string{
		return fmt.Sprintf("recived -> %s", t)
	} ("kakakka")

	fmt.Println(r)
}