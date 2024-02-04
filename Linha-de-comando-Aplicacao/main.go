package main

import (
	"fmt"
	"log"
	"os"
	"udemy/Linha-de-comando-Aplicacao/app"
)

func main() {
	fmt.Println("hi")

	//a aplicacao chama o pacote app e chama a funcao Gerar()
	// a variavel aplicacao é do tipo cli.App
	aplicacao := app.Gerar()
	//os.Args parametro padrao para que os comando do sistema operaciona seja reconhecido pela linha de comando
	//metodo Run pode retornar um erro, precisa tratar
	erro := aplicacao.Run(os.Args)
	if erro != nil {
		log.Fatal(erro) //parecido com fmt, mas log vem mais info e Fatal faz o programa parar
	}
}

