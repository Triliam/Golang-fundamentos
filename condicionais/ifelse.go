package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("numero maior que 15")
	} else {
		fmt.Println("numero menor que 15")
	}

	if novoNumero := numero; novoNumero > 1 {
		fmt.Println("Número é maior que 1")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

	if outroNumero := 5; outroNumero >=5 {
		fmt.Println("eh maior ou igual a 5")
	}
}