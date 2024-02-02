//funcoes que referenciam variaveis que tao FORA do seu corpo/escopo

//como funcoes tbm sao um tipo em GO podemos usar como parametro.. retorno...

package main

import "fmt"

func closure() func() {
	texto := "Dentro da funcao closure"

	//criar uma funcao sem retorno e sem parametro

	funcao := func ()  {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro do main"
	fmt.Println(texto)

	//variavel funcaoNova acaba sendo uma funcao, pois recebe o valor "closure()", onde o retorno é uma funcao
	funcaoNova := closure()
	funcaoNova()
}
//mantem a referencia inicial do valor da variavel, mesmo se usar uma outra variavel com o mesmo nome. No caso a variavel "texto"