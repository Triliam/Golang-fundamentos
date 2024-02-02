package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá World!")
	// chama a funcao com visibilidade Publica (inicial Maiuscula) Escrever(), e Escrever() chama dentro de si escrever() 
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("tais@gmail.com")
	erro2 := checkmail.ValidateFormat("dsadsad")
	fmt.Println(erro)
	fmt.Println(erro2)
}

// go mod tidy (limpa no go.mod  dependencias/bibliotecas que nao estao sendo usadas)