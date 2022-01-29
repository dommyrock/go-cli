package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/TwiN/go-color"
	"github.com/urfave/cli/v2"
)

func main() {

	inputFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app := &cli.App{
		Name:    "DNS lookup Tool",
		Usage:   "Provides More info about DNS. \nLets you query IP's, CNAME's,MX (mail exchange) records and Name Servers",
		Version: "0.0.1",
		Commands: []*cli.Command{
			{
				Name:  "ns",
				Usage: "Look name servers for particular Host",
				Flags: inputFlags,
				Action: func(c *cli.Context) error {
					fmt.Println(color.Colorize(color.Cyan, "DNS NS (name servers) records for the given domain name:"))

					ns, er := net.LookupNS(c.String("host"))
					if er != nil {
						fmt.Println(er)
						return er
					}
					for i := 0; i < len(ns); i++ {
						fmt.Println(ns[i].Host)
					}
					return nil
				},
			},
			{
				Name:  "ip",
				Usage: "Look IP's for particular Host",
				Flags: inputFlags,
				Action: func(c *cli.Context) error {
					fmt.Println(color.Colorize(color.Cyan, "IP's for specified Host:"))

					ip, er := net.LookupIP(c.String("host"))
					if er != nil {
						fmt.Println(er)
						return er
					}
					for i := 0; i < len(ip); i++ {
						fmt.Println(ip[i])
					}
					return nil
				},
			},
			{
				Name:  "cname",
				Usage: "Look CNAME's for particular Host",
				Flags: inputFlags,
				Action: func(c *cli.Context) error {
					fmt.Println(color.Colorize(color.Cyan, "CNAME for specified Host:"))

					cname, er := net.LookupCNAME(c.String("host"))
					if er != nil {
						fmt.Println(er)
						return er
					}
					fmt.Println(cname)
					return nil
				},
			},
			{
				Name:  "mx",
				Usage: "Look MX (mail exchange) records for particular Host",
				Flags: inputFlags,
				Action: func(c *cli.Context) error {
					fmt.Println(color.Colorize(color.Cyan, "MX records for specified Host:"))

					mx, er := net.LookupMX(c.String("host"))
					if er != nil {
						fmt.Println(er)
						return er
					}
					for i := 0; i < len(mx); i++ {
						fmt.Printf("pref:%d | Host: %s\n", mx[i].Pref, mx[i].Host)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
