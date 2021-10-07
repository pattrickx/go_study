package main

import "fmt"

type user struct {
	name string
	age uint8
	addr address
}
type address struct{
	street string
	number uint16
}

func main()  {
	fmt.Println("struct")

	var u user
	u.name  = "kjkjkj"
	u.age =  23
	fmt.Println(u)

	addr1 := address{"sasas",1600}
	u2 := user{"user2",45,addr1}
	fmt.Println(u2)

	u3 := user{age:36}
	u3.name ="user3"
	fmt.Println(u3)
	
}