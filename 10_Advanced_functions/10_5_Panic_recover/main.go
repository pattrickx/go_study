package main

import "fmt"

func recover_run(){
	if r:= recover(); r != nil{	
	fmt.Println("recover run ok")
	}
}

func student_note(n1,n2 float64) bool{
	defer recover_run()
	mean := (n1+n2)/2
	if mean>6{
		return true
	}else if mean <6{
		return false
	}

	panic("Note is 6")
}


func main()  {
	fmt.Println(student_note(6,6))
	fmt.Println("after run")
}