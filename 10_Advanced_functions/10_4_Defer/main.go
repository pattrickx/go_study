package main

import "fmt"

func func1()  {
	fmt.Println("lklklkl")
}

func func2()  {
	fmt.Println("kkkkkkkkkk")
}

func main()  {
	defer func1()
	fmt.Println("aaaaaaaaaaaaaaa")
	func2()
}