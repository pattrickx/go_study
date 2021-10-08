package main

import "fmt"

func calc(n1,n2 int32)(sum int32, subtraction int32){
	sum = n1+n2
	subtraction = n1-n2
	return 
}
func main()  {
	fmt.Println("Named_return_parameters")

	sum, sub := calc(10,15)
	fmt.Println(sum,sub)

}