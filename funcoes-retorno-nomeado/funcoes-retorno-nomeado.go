package main

import "fmt"

func calc(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {

	calculo1, calculo2 := calc(10, 20) 
		fmt.Println(calculo1, calculo2)
	
}