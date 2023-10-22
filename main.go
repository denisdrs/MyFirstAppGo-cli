package main

import (
	"discover/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting...")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
