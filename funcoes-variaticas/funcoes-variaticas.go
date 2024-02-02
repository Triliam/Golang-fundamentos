package main

import "fmt"

func calc(numeros...int) int {
	total := 0
	for _, numero := range numeros {
		total = total + numero
	}
	return total
}

func escrever(texto string, numeros...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

//pode ter apenas um parametro variatico por funcao: (texto string, outroTexto...string, num...int) n conseguiria colocar o num...int, e o variatico tem que ser o ultimo parametro da funcao
func gratz(texto string, outroTexto...string) {
	for _, ot := range outroTexto {
		fmt.Println(texto, ot)
	}
}


func main() {
	soma := calc(10, 11, 20, 30)
	fmt.Println(soma)

	escrever("Hi", 1,2,3,4)

	gratz("Hi", "Lorena", "Mafalda", "Jonathan")
}