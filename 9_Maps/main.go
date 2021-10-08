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

	user2 := map[string]map[string]string{
		"Nome": {
		"first":"Pedro",
		"last": "lala"},
		
		"id":{"number":"12"},
	}
	fmt.Println(user2)	
	delete(user2,"id")
	fmt.Println(user2)	
	user2["street"]= map[string]string{"number":"54"}
	fmt.Println(user2)	
}