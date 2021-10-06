package main

import "fmt"

func sum(n1 int32,n2 int32)int32  {
	return n1+n2
}

func calc(n1,n2 int32)(int32,int32){

	return sum(n1,n2),n1-n2
}

func main()  {

	sum1 := sum(10,20)
	fmt.Println(sum1)

	var s = func(txt string)string{
		fmt.Println(txt)
		return txt
	}

	result := s("blablabla")

	fmt.Println(result)

	fmt.Println(calc(12,23))
	
}