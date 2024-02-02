package main

import "fmt"

func funcao1(){
	fmt.Println("funcao 1")
}

func funcao2(){
	fmt.Println("funcao 2")
}

func media(n1, n2 float32) bool {
	//aqui o defer é a ultima coisa que retorna antes do return 
	defer fmt.Println("Média calculada. O resultado será retornado:")
	fmt.Println("Entrando na função para calcular media")
	resultado := (n1 + n2)/ 2
	
	if resultado >= 6 {
		return true
	} else {
		return false
	}
	
}

func main() {
	// DEFER = ADIAR, adia a execução da funcão ate o ultimo momento possivel
	defer funcao1()
	funcao2()

	fmt.Println(media(6, 6))
}