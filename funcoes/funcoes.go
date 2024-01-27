package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Pode ter mais de um retorno e serem de tipos diferentes
func calculos(txt string, n1 int8, n2 int8) (string, int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return txt, soma, sub // se escrever algo no lugar de txt, mantem esse texto ao inves do texto do parametro da funcao quando chama a funcao
}

func calculosM(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub 
}

func main() {

	// pode criar uma função e atribuir ela a uma variavel
	var variavelDeUmaFuncao = func(t string) string {
		fmt.Println("sahuisa")
		return t
	}
	// dai chamar a variavel (que tem como valor a funcao) e conseguir o resultado da funcao atribuida a ela
	resultadoTexto := variavelDeUmaFuncao("aqui vai o texto")
	fmt.Println(resultadoTexto)

	soma := somar(10, 15)
	fmt.Println(soma)

	resultadoString, resultadoSoma, resultadoSub := calculos("o texto", 10, 15)
	fmt.Println(resultadoString, resultadoSoma, resultadoSub)

	rSoma, rSub := calculosM(1, 1) 
	fmt.Println(rSoma, rSub)

	// pode escolher quais retornos quer e anular outros
	_, rSub1 := calculosM(1, 1) 
	fmt.Println(rSub1)

	resultadoString1, _, resultadoSub1 := calculos("o texto", 10, 15)
	fmt.Println(resultadoString1, resultadoSub1)

}