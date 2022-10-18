package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "host",
		Value: "devbook.com.br",
	},
}

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicaco de linha de comando"
	app.Usage = "Busca IP's e nomes de servidor na internet"

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IP's de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Buscar servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
