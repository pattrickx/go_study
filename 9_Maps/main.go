package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Maps")

	user := map[string]string{
		"Nome": "Pedro",
		"Id": "12",
	}
	fmt.Println(user)

	user2 := map[string]interface{}{
		"Nome": map[string]string{
		"first":"Pedro",
		"last": "lala"},
		
		"id":12,
	}
	fmt.Println(user2)	
	delete(user2,"id")
	fmt.Println(user2)	
	user2["street"]= map[string]string{"number":"54"}
	fmt.Println(user2)	
}