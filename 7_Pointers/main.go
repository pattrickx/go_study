package main

import "fmt"

func main()  {
	fmt.Println("pointers")
	
	var v1 int = 10
	var v2 int = v1
	fmt.Println(v1,v2)

	v1++
	fmt.Println(v1,v2)

	//Pointers
	var v3 int = 10
	var p1 *int = &v3

	fmt.Println(v3,p1)
	fmt.Println(v3,*p1)

	v3++
	fmt.Println(v3,*p1)

}