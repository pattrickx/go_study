package main

import (
	"fmt"
	"reflect"
)

func main()  {
	fmt.Println("8_Arrays_Slices")

	var a1 [5]int
	a1[0]=1
	fmt.Println(a1)

	a2 := [5]string{"s","kkjk","kjkjkj","klklkl", "pppoo"}
	fmt.Println(a2)

	a3 := [...]int{1,2,3,4,5}
	fmt.Println(a3)

	slice := []int{1,2,5,7,6,8,9,71,1}
	fmt.Println(slice)
	slice = append(slice, 5)
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(a3))
	fmt.Println(reflect.TypeOf(slice))

	slice1 := a2[1:3]
	fmt.Println(slice1)
	a2[2] = "jaja"
	fmt.Println(slice1)

	// inter arrays
	fmt.Println("---------")
	slice3 := make([]float32,10,11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // size
	fmt.Println(cap(slice3)) // max size

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // size
	fmt.Println(cap(slice3)) // max size

	slice3 = append(slice3, 8)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // size
	fmt.Println(cap(slice3)) // max size

	slice4 := make([]float32,5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // size
	fmt.Println(cap(slice4)) // max size


}