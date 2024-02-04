package main

import (
	"fmt"
	"sync"
) 

func escrever(txt string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(txt)
	}
}

func main() {
	//cria uma variavel do tipo waitgroup
	//pacote.Tipo
	var variavelDoTipoWaitGroup sync.WaitGroup

	//diz pro waitgroup quantas gorotines ele vai lidar
	variavelDoTipoWaitGroup.Add(3)

	go func(){
		escrever("funcao 1")
		variavelDoTipoWaitGroup.Done()
	}()

	go func(){
		escrever("funcao 2")
		variavelDoTipoWaitGroup.Done()
	}()

	go func(){
		escrever("funcao3")
		variavelDoTipoWaitGroup.Done()
	}()

	variavelDoTipoWaitGroup.Wait()

	fmt.Println("terminou")
}
//qnd n quer que uma funcao dependa da outra terminar para executar: Sincronizar as gorotines, onde as duas funcoes terminem antes do programa terminar de executar