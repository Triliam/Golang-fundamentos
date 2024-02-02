//usada para configurações iniciais, ele inicializa antes da funcao main

package main

import "fmt"

//declara fora
var txt string
var n int

func init(){
	fmt.Println("funcao init")
	n = 10
	txt = "config iniciais"
}

func main() {
	fmt.Println("funcao main")
	fmt.Println(txt)
	fmt.Println(n)
}