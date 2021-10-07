package main

import "fmt"

type people struct{
	name string
	age uint8
}

type student struct{
	people
	course string
	college string
}

func main()  {
	p1 := people{"kjkj",35}
	s1 := student{p1,"eng","f1"}
	fmt.Println(s1)
	fmt.Println(s1.name)
}