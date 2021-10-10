package main

import "fmt"

func generic(ineterf interface{}){
	fmt.Println(ineterf)
}

func main()  {
	fmt.Println("generic type")
	generic(10)
	generic('k')
	generic("sadas")
	generic(10.5)
}