package main

import "fmt"


func main() {

	func() {
		fmt.Println("Hi mundu")
	}()

	func(texto string) {
		fmt.Println("Hi mundu", texto)
	}("aqui o texto")

	variavelDoRetorno := func(txt string) string {
		return fmt.Sprintf("um texto e %s", txt)
	}("outro texto")

	fmt.Println(variavelDoRetorno)
}