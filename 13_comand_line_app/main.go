package main

import (
	"comand_line_app/app"
	"fmt"
	"log"
	"os"
)



func main()  {
	fmt.Println("comand line app")

	app1 := app.Generate() 
	
	if erro := app1.Run(os.Args); erro != nil{
		log.Fatal(erro)
	}
}