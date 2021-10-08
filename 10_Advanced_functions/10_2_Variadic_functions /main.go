package main

import "fmt"

func sum(nums ...int)int{
	s := 0
	for _,n := range nums{
		s+=n

	}
	return s
} 
func show(s string, nums ...int){
	
		fmt.Println(s)
		fmt.Println(nums)
}
func main()  {
	fmt.Println("10_2_Variadic_functions ")


	fmt.Println(sum(1,2,35,8,9,7))
	show("kjkj",1,2,35,8,9,7)

}