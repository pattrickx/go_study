package main

import "fmt"

type user struct{
	name string
	age uint8
}

func (u user) save(){
	fmt.Printf("saving data of user %s\n",u.name)
}

func (u user) adult() bool {
	return u.age >=18
}

func (u *user) birthday(){
	u.age ++
}
func main()  {
	fmt.Println("methods")

	user1 := user{"pppp",20}
	user1.save()
	fmt.Println(user1.adult())

	user1.birthday()
	fmt.Println(user1.age)



}