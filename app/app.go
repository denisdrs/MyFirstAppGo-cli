package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de CLI pronta para ser executada.
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Discover"
	app.Usage = "Busca IP e Nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "kernel.org",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Retorna IP do servidor na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "hostname",
			Usage:  "Busca nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app

}

// Implementa buscar ips
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

	servidores, erro := net.LookupNS(host) // name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
