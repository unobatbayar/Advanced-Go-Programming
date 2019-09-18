package main

import (
	"log"
	"os"
	"net"
	"fmt"

	"github.com/urfave/cli"
)

func main(){
	app := cli.NewApp()
	app.Name = "Command Line Interface Tool"
	app.Usage = "Use for name server lookup, IP lookup and more."

	flags := []cli.Flag{
		cli.StringFlag{ 
			Name: "host",
			Value: "konple.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ns",
			Usage: "Looks up the name server for the given parameter",
			Flags: flags,
			Action: func(c *cli.Context) error{
				ns, err := net.LookupNS(c.String("host"))
				if err != nil{
					return err
				}
				for i := 0; i < len(ns); i++{
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},

		{
			Name: "ip",
			Usage: "Looks up the IP address for the given paramter",
			Flags: flags,
			Action: func(c *cli.Context) error{
				ip, err := net.LookupIP(c.String("host"))
				if err != nil{
					fmt.Println(err)
				}
				for i := 0; i < len(ip); i++{
					fmt.Println(ip[i])
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}