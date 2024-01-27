package main


import (
	"fmt"
)

func main() {
	var variavel1 string = "variavel 1"
	//inferencia de tipo - a partir do valor da variavel
	variavel2 := "variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "hadsa"
		variavel4 string = "hsgadgdsa"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "valor 5", "valor 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "contante 1"

	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}