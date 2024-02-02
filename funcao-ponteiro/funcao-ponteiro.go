//as funcoes podem usar um ponteiro como parametro ou retorno ao inves de uma variavel

package main

import "fmt"

//n altera o valor da variavel
//copia
func inverterSinal(numero int) int {
	return numero * -1
} 
// altera o valor da variavel toda vez que chamar a funcao
//referencia
func inverterSinalComPonteiro(numero *int) {
	// usa o * pra buscar o valor que esta no endereço de memoria
	*numero = *numero *-1
}

func main() {
	n := 20 // n recebe 20

	numeroInvertido := inverterSinal(n) // numeroInvertido vai receber -20 
	fmt.Println(numeroInvertido) //vale/armazena -20
	fmt.Println(n) // n continua valendo/armazenando 20

	novoNumero := 40
	fmt.Println(novoNumero)
	// usa o & para trazer pra variavel o endereço de memoria onde o ponteiro aponta
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}