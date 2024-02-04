//canais de comunicacao tem duas operaçoes: receber e enviar dados para sincronizar as gorotines
//essas operacoes bloqueiam a execucao do programa
//mesmo rodando uma concorrencia, por ser uma operacao BLOQUEANTE o canal espera até receber o dado. É usada em gorotines onde vc PRECISA esperar receber um valor para poder continuar o programa
//cuidar para n dar deadlock se o canal estiver esperando receber um dado e nao tiver mais dado pra ser eviado
package main

import (
	"fmt"
)

func escrever(txt string, canal chan string) {

	for i := 0; i<= 5; i++ {
		canal <- txt
	}
close(canal)	
}

func main() {
	//criando um canal, no caso do tipo string

	canal := make(chan string)
	go escrever("texto que o canal vai receber", canal)

	fmt.Println("linha depois da funcao escrever iniciar a execussão")

	// for {
	// 	mensagem, verificaCanal := <- canal
	// 	if !verificaCanal {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	// pra cada mensagem que for recebida no canal enquanto ele estiver aberto, printa a msg
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("fim")
}