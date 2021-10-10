package main

import (
	"fmt"
	"math"
)


type ret struct{
	w float64
	h float64
}
func (r ret) area() float64{
	return r.h*r.w
}

type circle struct{
	r float64
}
func (c circle) area() float64{
	return math.Pi * math.Pow(c.r,2)
}
func area(s shape){
	fmt.Printf("area is %0.2f\n",s.area())
}
type shape interface{
	area() float64

}

func main()  {
	fmt.Println("Interfaces")
	r := ret{10,15}
	c := circle{4}
	area(r)
	area(c)
	
}