package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// return comand line app
func Generate() *cli.App {
	app:= cli.NewApp()
	app.Name = "Comand line app"
	app.Usage = "Search IPS and server names on the internet"

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Search IPS and server names on the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "github.com",
				},
			},
			Action: shearchIps,
		},
	}
	return app
}

func shearchIps(c *cli.Context)  {
	host := c.String("host")

	ips,erro := net.LookupIP(host)
	if erro != nil{
		log.Fatal(erro)
	}
	for _,ip:= range ips{
		fmt.Println(ip)
	}
}